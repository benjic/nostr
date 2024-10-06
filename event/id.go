package event

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type ID [32]byte

func (p ID) MarshalJSON() ([]byte, error) { return []byte(fmt.Sprintf(`"%032x"`, p)), nil }

func newID(p Payload, pk PubKey) ID {
	bs, _ := json.Marshal(
		[]any{0, pk, p.Timestamp, p.Kind, p.Tags, p.Content},
	)

	return ID(sha256.Sum256(bs))
}
