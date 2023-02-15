package operations

import (
	"openapi/pkg/models/shared"
	"openapi/pkg/models/utils"
)

type DeletePetPathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePetRequest struct {
	PathParams DeletePetPathParams
	Retries    *utils.RetryConfig
}

type DeletePetResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
