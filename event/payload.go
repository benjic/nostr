package event

type Kind int16

type Payload struct {
	Timestamp int64  `json:"created_at"`
	Kind      Kind   `json:"kind"`
	Tags      Tags   `json:"tags"`
	Content   string `json:"content"`
}
