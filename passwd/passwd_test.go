package passwd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Example = "gleb:x:1000:1001::/home/gleb:/usr/bin/fish"

func TestParse(t *testing.T) {
	users, err := Parse(Example)

	assert.Equal(t, nil, err, "should be no error")
	assert.Equal(t, 1, len(users), "users should have one")

	assert.Equal(t, "gleb", users[0].Name, "")
	assert.Equal(t, "1000", users[0].UID, "")
	assert.Equal(t, "1001", users[0].GID, "")
	assert.Equal(t, "", users[0].Desc, "")
	assert.Equal(t, "/home/gleb", users[0].Home, "")
	assert.Equal(t, "/usr/bin/fish", users[0].Shell, "")

	assert.Equal(t, Example, users[0].Serialize(), "")

	assert.Equal(t, "useradd -m -u 1000 -g 1001 -p x -d /home/gleb -s /usr/bin/fish gleb", users[0].UserAdd(), "")
}
