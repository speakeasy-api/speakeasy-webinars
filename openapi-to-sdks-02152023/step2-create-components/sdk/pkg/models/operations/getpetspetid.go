package operations

import (
	"openapi/pkg/models/shared"
)

type GetPetsPetIDPathParams struct {
	PetID string `pathParam:"style=simple,explode=false,name=petId"`
}

type GetPetsPetIDRequest struct {
	PathParams GetPetsPetIDPathParams
}

type GetPetsPetIDResponse struct {
	ContentType string
	Error       *shared.Error
	Pet         *shared.Pet
	StatusCode  int64
}
