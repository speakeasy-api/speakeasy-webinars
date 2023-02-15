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
    
    req := operations.AddPetRequest{
        Request: shared.NewPet{
            Name: "unde",
            Tag: "deserunt",
        },
    }
    
    res, err := s.Pets.AddPet(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->