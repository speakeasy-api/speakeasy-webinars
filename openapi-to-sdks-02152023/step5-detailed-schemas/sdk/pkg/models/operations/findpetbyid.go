package operations

import (
	"openapi/pkg/models/shared"
)

type FindPetByIDPathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindPetByIDRequest struct {
	PathParams FindPetByIDPathParams
}

type FindPetByIDResponse struct {
	ContentType string
	Error       *shared.Error
	Pet         *shared.Pet
	StatusCode  int64
}
