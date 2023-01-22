package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/a7vicky/admissioncontroller-demo/pkg/serve"
	log "k8s.io/klog/v2"
)

var (
	listenAddress = flag.String("listen", "0.0.0.0", "listen address")
	listenPort    = flag.String("port", "8443", "port to listen on")

	useTLS  = flag.Bool("tls", false, "Use TLS? Must specify -tlskey, -tlscert, -cacert")
	tlsKey  = flag.String("tlskey", "", "TLS Key for TLS")
	tlsCert = flag.String("tlscert", "", "TLS Certificate")
	caCert  = flag.String("cacert", "", "CA Cert file")
)

func main() {

	flag.Parse()

	server := serve.NewServer(*listenAddress, *listenPort)

	go func() {
		// listen shutdown signal
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		sig := <-signalChan
		log.Errorf("Received %s signal; shutting down...", sig)
		if err := server.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
	}()

	if *useTLS {
		cafile, err := ioutil.ReadFile(*caCert)
		if err != nil {
			log.Error(err, "Couldn't read CA cert file")
			os.Exit(1)
		}
		certpool := x509.NewCertPool()
		certpool.AppendCertsFromPEM(cafile)

		server.TLSConfig = &tls.Config{
			RootCAs: certpool,
		}
		log.Error(server.ListenAndServeTLS(*tlsCert, *tlsKey), "Error serving TLS")
		log.Infof("Starting server on port: %s", *listenPort)
	} else {
		log.Error(server.ListenAndServe(), "Error serving non-TLS connection")
	}
}
