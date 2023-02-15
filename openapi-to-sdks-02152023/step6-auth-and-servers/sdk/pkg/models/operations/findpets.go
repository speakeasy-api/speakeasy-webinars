package operations

import (
	"openapi/pkg/models/shared"
)

type FindPetsQueryParams struct {
	Limit *int32   `queryParam:"style=form,explode=true,name=limit"`
	Tags  []string `queryParam:"style=form,explode=true,name=tags"`
}

type FindPetsRequest struct {
	QueryParams FindPetsQueryParams
}

type FindPetsResponse struct {
	ContentType string
	Error       *shared.Error
	Pets        []shared.Pet
	StatusCode  int64
}
