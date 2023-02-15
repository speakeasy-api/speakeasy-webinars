package shared

type NewPet struct {
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}
