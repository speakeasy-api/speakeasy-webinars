package operations

import (
	"openapi/pkg/models/shared"
)

type PostPetsResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
