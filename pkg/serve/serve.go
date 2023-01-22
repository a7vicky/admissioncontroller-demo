package serve

import (
	"fmt"
	"net/http"

	"github.com/a7vicky/admissioncontroller-demo/pkg/admissionhandler"
	"github.com/a7vicky/admissioncontroller-demo/pkg/validation"
)

func NewServer(address string, port string) *http.Server {

	// instance of hook
	nsValidation := validation.NewValidationHook()

	// router
	ah := admissionhandler.NewAdmissionHandler()
	mux := http.NewServeMux()
	mux.Handle("/healthz", admissionhandler.Healthz())
	mux.HandleFunc("/validate-ns", ah.Serve(nsValidation))

	return &http.Server{
		Addr:    fmt.Sprintf(address, port),
		Handler: mux,
	}
}
