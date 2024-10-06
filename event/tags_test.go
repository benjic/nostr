package event

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Tags_MarshalJSON(t *testing.T) {
	bs, err := json.Marshal(Tags(nil))
	require.NoError(t, err)
	assert.Equal(t, "[]", string(bs))
}

func Test_Tag_MarshalJSON(t *testing.T) {
	bs, err := json.Marshal(Tag(nil))
	require.NoError(t, err)
	assert.Equal(t, "[]", string(bs))
}
