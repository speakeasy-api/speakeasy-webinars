package operations

import (
	"openapi/pkg/models/shared"
)

type DeletePetPathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePetRequest struct {
	PathParams DeletePetPathParams
}

type DeletePetResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
