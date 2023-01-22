# admissioncontroller-demo

### genrate-certificate

$ openssl req -newkey rsa:2048 -nodes -keyout tls.key -out tls.csr
$ openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost,DNS:127.0.0.1") -days 365 -in tls.csr -signkey tls.key -out tls.crt
Encode tls.crt
$ cat tls.crt | base64 | tr -d " \t\n\r"