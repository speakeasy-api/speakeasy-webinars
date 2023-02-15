<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "openapi"
    "openapi/pkg/models/shared"
    "openapi/pkg/models/operations"
)

func main() {
    s := sdk.New()
    
    req := operations.GetPetsRequest{
        QueryParams: operations.GetPetsQueryParams{
            Limit: 548814,
        },
    }
    
    res, err := s.GetPets(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->