package event

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

var (
	ErrInvalidSignature = fmt.Errorf("invalid signature")
	ErrInvalidID        = fmt.Errorf("invalid id")
)

func Validate(e Event) error {
	pk, err := schnorr.ParsePubKey(e.PubKey[:])
	if err != nil {
		return fmt.Errorf("%w: failed to parse public key: %w", ErrInvalidSignature, err)
	}

	sig, err := schnorr.ParseSignature(e.Signature[:])
	if err != nil {
		return fmt.Errorf("%w: failed to parse signature: %w", ErrInvalidSignature, err)
	}

	if !sig.Verify(e.ID[:], pk) {
		return fmt.Errorf("%w: failed to verify signature", ErrInvalidSignature)
	}

	if e.ID != newID(e.Payload, e.PubKey) {
		return fmt.Errorf("%w: %w", ErrInvalidSignature, ErrInvalidID)
	}

	return nil
}
