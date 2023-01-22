package validation

import (
	"github.com/a7vicky/admissioncontroller-demo/pkg/admissionwebhook"

	"k8s.io/api/admission/v1beta1"
)

func validateCreate() admissionwebhook.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*admissionwebhook.Result, error) {
		namespace, err := parseNamespace(r.Object.Raw)
		if err != nil {
			return &admissionwebhook.Result{Msg: err.Error()}, nil
		}
		if namespace.Name == "demo" {
			return &admissionwebhook.Result{Msg: "You cannot create namespace 'demo'."}, nil
		}
		return &admissionwebhook.Result{Allowed: true}, nil
	}
}
