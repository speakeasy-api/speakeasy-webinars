package webhooks

import (
	"openapi/pkg/models/shared"
)

type NewPetResponse struct {
	ContentType string
	StatusCode  int64
}

type NewPetRequest struct {
	Request *shared.NewPet `request:"mediaType=application/json"`
}
