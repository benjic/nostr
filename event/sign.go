package event

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

type PubKey [32]byte

func (p PubKey) MarshalJSON() ([]byte, error) { return []byte(fmt.Sprintf(`"%032x"`, p)), nil }

type Signature [64]byte

func (p Signature) MarshalJSON() ([]byte, error) { return []byte(fmt.Sprintf(`"%064x"`, p)), nil }

type Event struct {
	Payload

	ID        ID        `json:"id"`
	PubKey    PubKey    `json:"pubkey"`
	Signature Signature `json:"sig"`
}

func Sign(p Payload, key []byte) (Event, error) {
	sk, pk := btcec.PrivKeyFromBytes(key)
	pubKey := PubKey(pk.SerializeCompressed()[1:])
	id := newID(p, pubKey)

	sig, err := schnorr.Sign(sk, []byte(id[0:32]))
	if err != nil {
		return Event{}, fmt.Errorf("failed to sign event: %w", err)
	}

	return Event{Payload: p, ID: id, PubKey: pubKey, Signature: Signature(sig.Serialize())}, nil
}
