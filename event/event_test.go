package event

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testPrivateKey = []byte{
		0x2e, 0x42, 0xf5, 0x6, 0x80, 0xe6, 0x22, 0x8b,
		0x8d, 0xf7, 0xa4, 0x84, 0x40, 0x51, 0x8a, 0x35,
		0x57, 0x19, 0xdf, 0x22, 0x3a, 0xaf, 0x5, 0xf2,
		0x98, 0x69, 0x6b, 0x8c, 0xc3, 0x86, 0x29, 0x2f,
	}
)

func Test_MarshalJSON(t *testing.T) {
	t.Run("from builder", func(t *testing.T) {
		b := NewBuilder(1, "Hello World", testPrivateKey)

		ts, err := time.Parse(time.DateOnly, "2024-10-05")
		require.NoError(t, err)
		b.SetTimestamp(ts)

		e, err := b.Build()
		require.NoError(t, err)

		assert.NoError(t, Validate(e))

		bs, err := json.Marshal(e)
		require.NoError(t, err)

		assert.Equal(
			t,
			`{"created_at":1728086400,"kind":1,"tags":[],"content":"Hello World","id":"bab012190e329f702eedcf169f8fd0ecb3a3a6e4fd397fb52e7a0c762fb75cd3","pubkey":"e64c94c839b9b8ef4b6c949dc1d18dc46d2ecd7cfd2d42ec66e86ba3541819b6","sig":"cf1768f3465d0b3057f37cab34e6c8caa1cf777c25ca17489757cba3bb38f205944777d8754d748e46504acdadcfea97b7ee526d9dc3fa33d77586191e4bf460"}`,
			string(bs),
		)
	})

	t.Run("from primitives", func(t *testing.T) {
		t.Run("empty value", func(t *testing.T) {
			e, err := Sign(Payload{}, testPrivateKey)
			assert.NoError(t, err)

			assert.NoError(t, Validate(e))

			bs, err := json.Marshal(e)
			require.NoError(t, err)

			assert.Equal(
				t,
				`{"created_at":0,"kind":0,"tags":[],"content":"","id":"9937133004cf0e0759a2bbc11ae7541ec597d7e0e09c78ef2c40c70dcc5682a3","pubkey":"e64c94c839b9b8ef4b6c949dc1d18dc46d2ecd7cfd2d42ec66e86ba3541819b6","sig":"8dd20d35f01f34b4851f3045968200f0d1905f74f691f86233ecd1236c176036ecd6fbe1c1897be093f893adc6f6d8809baa0ed9221154505c3fcd66dd1f8060"}`,
				string(bs),
			)
		})

		t.Run("given value", func(t *testing.T) {
			e, err := Sign(
				Payload{
					Timestamp: 105,
					Kind:      202,
					Tags: []Tag{
						{"e", "5c83da77af1dec6d7289834998ad7aafbd9e2191396d75ec3cc27f5a77226f36", "wss://nostr.example.com"},
						{"p", "f7234bd4c1394dda46d09f35bd384dd30cc552ad5541990f98844fb06676e9ca"},
					},
					Content: "Not empty content",
				},
				testPrivateKey,
			)
			assert.NoError(t, err)

			assert.NoError(t, Validate(e))

			bs, err := json.Marshal(e)
			require.NoError(t, err)

			assert.Equal(
				t,
				`{"created_at":105,"kind":202,"tags":[["e","5c83da77af1dec6d7289834998ad7aafbd9e2191396d75ec3cc27f5a77226f36","wss://nostr.example.com"],["p","f7234bd4c1394dda46d09f35bd384dd30cc552ad5541990f98844fb06676e9ca"]],"content":"Not empty content","id":"3afa2c8747d62494ea26a1b94978cba1737ee93020eb3950e2642160e0ad0691","pubkey":"e64c94c839b9b8ef4b6c949dc1d18dc46d2ecd7cfd2d42ec66e86ba3541819b6","sig":"71de4bd57d91532dd48f78a99972accd3793d8d112b95c4eea87f405541756801451ad6e9dfacaa6b9673161dd40ac7e43a935d5c70dc43a1c4decab90a78214"}`,
				string(bs),
			)
		})
	})
}
