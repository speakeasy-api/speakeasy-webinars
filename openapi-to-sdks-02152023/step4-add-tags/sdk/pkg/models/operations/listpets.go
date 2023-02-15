package operations

import (
	"openapi/pkg/models/shared"
)

type ListPetsQueryParams struct {
	Limit *int32 `queryParam:"style=form,explode=true,name=limit"`
}

type ListPetsRequest struct {
	QueryParams ListPetsQueryParams
}

type ListPetsResponse struct {
	ContentType string
	Error       *shared.Error
	Pets        []shared.Pet
	StatusCode  int64
}
