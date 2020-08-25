package group

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Example = "gleb:x:1000:test1,test2"

func TestParse(t *testing.T) {
	groups, err := Parse(Example)

	assert.Equal(t, nil, err, "should be no error")
	assert.Equal(t, 1, len(groups), "groups should have one")

	assert.Equal(t, "gleb", groups[0].Name, "")
	assert.Equal(t, "x", groups[0].Password, "")
	assert.Equal(t, "1000", groups[0].GID, "")
	assert.Equal(t, "test1,test2", groups[0].Members, "")

	assert.Equal(t, Example, groups[0].Serialize(), "")

	assert.Equal(t, "groupadd -g 1000 gleb", groups[0].GroupAdd(), "")
}
