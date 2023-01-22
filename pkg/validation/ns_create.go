package validation

import (
	"github.com/a7vicky/admissioncontroller-demo/pkg/admissionwebhook"
	admission "k8s.io/api/admission/v1"
)

func validateCreate() admissionwebhook.AdmitFunc {
	return func(r *admission.AdmissionRequest) (*admissionwebhook.Result, error) {
		namespace, err := parseNamespace(r.Object.Raw)
		if err != nil {
			return &admissionwebhook.Result{Msg: err.Error()}, nil
		}
		if namespace.Name == "demo" {
			return &admissionwebhook.Result{Allowed: false}, nil
		}
		return &admissionwebhook.Result{Allowed: true}, nil
	}
}
