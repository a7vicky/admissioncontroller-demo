# admissioncontroller-demo

### generate-certificate

```bash
$ openssl req -newkey rsa:2048 -nodes -keyout tls.key -out tls.csr
$ openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost,DNS:127.0.0.1,DNS:0.0.0.0") -days 365 -in tls.csr -signkey tls.key -out tls.crt
Encode tls.crt
$ cat tls.crt | base64 | tr -d " \t\n\r"
```

### Build and Start server

```bash
$ make build
$ ./bin/admission-webhook-demo -tls -tlskey hack/certs/tls.key -tlscert hack/certs/tls.crt -cacert hack/certs/tls.crt
```
### Test

```text
$ curl --cacert hack/certs/tls.crt https://localhost:8443/healthz
Create validationwebhook configuration
$ kubectl create -f deploy/webhook-example.yaml
$ kubectl create ns demo
```