package event

import (
	"fmt"
	"time"
)

var (
	ErrNoSignerSet = fmt.Errorf("no signer set for builder")
	ErrSignFailure = fmt.Errorf("failed to sign event")
)

type Builder struct {
	privKey []byte

	createdAt *time.Time
	kind      *Kind
	content   *string
	tags      *Tags
}

func NewBuilder(k Kind, c string, privKey []byte) *Builder {
	return &Builder{
		privKey: privKey,

		kind:    &k,
		content: &c,
	}
}

func (b *Builder) SetTags(t Tags) { b.tags = &t }

func (b *Builder) SetTimestamp(t time.Time) { b.createdAt = &t }

func (b *Builder) Build() (Event, error) {
	switch {
	case b.kind == nil, b.content == nil:
		return Event{}, fmt.Errorf("missing required data")
	}

	createdAt := time.Now()
	if b.createdAt != nil {
		createdAt = *b.createdAt
	}

	tags := Tags{}
	if b.tags != nil {
		tags = *b.tags
	}

	payload := Payload{
		Kind:      *b.kind,
		Content:   *b.content,
		Timestamp: createdAt.Unix(),
		Tags:      tags,
	}

	e, err := Sign(payload, b.privKey)
	if err != nil {
		return Event{}, fmt.Errorf("%w: %w", ErrSignFailure, err)
	}

	return e, nil
}
