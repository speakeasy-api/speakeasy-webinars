package operations

import (
	"openapi/pkg/models/shared"
)

type CreatePetResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
