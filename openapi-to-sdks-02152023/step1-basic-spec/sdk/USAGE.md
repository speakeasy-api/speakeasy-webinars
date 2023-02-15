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
    
    req := operations.PostPetRequest{
        Request: "unde",
    }
    
    res, err := s.PostPet(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostPet200ApplicationJSONAny != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->