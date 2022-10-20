package validators

import (
	"errors"
	"sur/application/messages"
)

func CreateUrlValidator(request messages.CreateUrlRequest) error {
	if request.TenantId == 0 {
		return errors.New("'Tenant' must not be empty")
	}
	if request.Url == "" {
		return errors.New("'Url' must not be empty")
	}
	return nil
}
