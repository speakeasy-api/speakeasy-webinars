package operations

import (
	"openapi/pkg/models/shared"
)

type GetPetPathParams struct {
	PetID string `pathParam:"style=simple,explode=false,name=petId"`
}

type GetPetRequest struct {
	PathParams GetPetPathParams
}

type GetPetResponse struct {
	ContentType string
	Error       *shared.Error
	Pet         *shared.Pet
	StatusCode  int64
}
