package operations

import (
	"openapi/pkg/models/shared"
)

type GetPetsQueryParams struct {
	Limit *int32 `queryParam:"style=form,explode=true,name=limit"`
}

type GetPetsRequest struct {
	QueryParams GetPetsQueryParams
}

type GetPetsResponse struct {
	ContentType string
	Error       *shared.Error
	Pets        []shared.Pet
	StatusCode  int64
}
