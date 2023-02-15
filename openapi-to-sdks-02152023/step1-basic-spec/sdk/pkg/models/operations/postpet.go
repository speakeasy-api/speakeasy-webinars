package operations

type PostPetRequest struct {
	Request *interface{} `request:"mediaType=application/json"`
}

type PostPetResponse struct {
	ContentType                  string
	StatusCode                   int64
	PostPet200ApplicationJSONAny *interface{}
}
