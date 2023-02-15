package operations

import (
	"openapi/pkg/models/shared"
	"openapi/pkg/models/utils"
)

type AddPetRequest struct {
	Request shared.NewPet `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type AddPetResponse struct {
	ContentType string
	Error       *shared.Error
	Pet         *shared.Pet
	StatusCode  int64
}
