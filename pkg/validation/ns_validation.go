package validation

import (
	"encoding/json"

	"github.com/a7vicky/admissioncontroller-demo/pkg/admissionwebhook"

	v1 "k8s.io/api/core/v1"
)

// NewValidationHook creates a new instance of namespace validation hook
func NewValidationHook() admissionwebhook.Hook {
	return admissionwebhook.Hook{
		Create: validateCreate(),
	}
}

func parseNamespace(object []byte) (*v1.Namespace, error) {
	var namespace v1.Namespace
	if err := json.Unmarshal(object, &namespace); err != nil {
		return nil, err
	}
	return &namespace, nil
}
