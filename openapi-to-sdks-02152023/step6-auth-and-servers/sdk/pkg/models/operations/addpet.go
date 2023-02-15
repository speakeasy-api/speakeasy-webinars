package operations

import (
	"openapi/pkg/models/shared"
)

type AddPetRequest struct {
	Request shared.NewPet `request:"mediaType=application/json"`
}

type AddPetResponse struct {
	ContentType string
	Error       *shared.Error
	Pet         *shared.Pet
	StatusCode  int64
}
