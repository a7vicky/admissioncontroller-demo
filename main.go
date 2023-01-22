package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/a7vicky/admissioncontroller-demo/pkg/admissionhandler"
	"github.com/a7vicky/admissioncontroller-demo/pkg/validation"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {

	// instance of hook
	nsValidation := validation.NewValidationHook()

	// router
	ah := admissionhandler.NewAdmissionHandler()
	mux := http.NewServeMux()
	mux.Handle("/healthz", admissionhandler.Healthz())
	mux.HandleFunc("/validate-ns", ah.Serve(nsValidation))

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
