package shared

type Pet struct {
	ID   int64   `json:"id"`
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}
