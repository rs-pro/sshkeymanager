package authorized_keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Example = "ssh-rsa AAAA4f test"

func TestParse(t *testing.T) {
	keys, err := Parse(Example)

	assert.Equal(t, nil, err, "should be no error")
	assert.Equal(t, 1, len(keys), "keys should have one")
	assert.Equal(t, "ssh-rsa AAAA4f", keys[0].Key, "")
	assert.Equal(t, "test", keys[0].Email, "")
}
