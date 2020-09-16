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

var Example2 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDC/yPQQ95IbMbDSw40Bf50mKSZSHfzOL3FH5zIcL1S2Ua7HaDGusEfZgibd4PBbLyKgRr9reXjtsv4aahIomJUhsYNiVYUyXU5fS0vG1ti2JXBWX11yr4QapXlgQBLLaeVS1RkzokIT3KPe3djxQ5YPDlA2sNPT5MsMyKYLEFIYDj3OC3RnoWmAS+EUB8ggT96Pa9Pj687MYeN3WXambfuts2m8Rp3KUAV37QBbmlZfpo4QKDvkVHApyHpP4RTlMPyYAuZfNw6RTPFyuELHARQph0mqbqOjRRhpKKpTDYAy4VrUTJ2q5OVuMeSMGBLAgVLxrvhEluaVO7wKdM1js+Yb96wElGhGzClbtOjDDHDZPXdstLUBnYVikIboQpbjhlwyekxLjADoOk/8LQrl8fHECR4DiVVm/ZPorgFa6E6iVJt/1gHtoJZpxxrQ1rjA3hwbaK7Pf+V5Fg2KGbioPjgKhjhSaLi9LVyd3ZX/7qUfKC3dNuaUNwQ3/4aCuDNoe9z6NCxSTI8V5sH3z0znddhZUiVKd21tTCuB1rdB+DN5xQKRenPbeWsow6g2v6GlNlNSEyzn5UjlLlHyBmZ82uZWJjgfa+OGvVvAo4R4Fk0ZvV8NMUhchEtl5e4DltNeiObKVzU81lUQvZNlJBH0twChYWiXwD+GoZHhDjr8b5xSw== Тест Юзер (rscz.ru)"

func TestParseFull(t *testing.T) {
	keys, err := Parse(Example2)
	assert.Equal(t, nil, err, "should be no error")
	assert.Equal(t, 1, len(keys), "keys should have one")
	assert.Equal(t, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDC/yPQQ95IbMbDSw40Bf50mKSZSHfzOL3FH5zIcL1S2Ua7HaDGusEfZgibd4PBbLyKgRr9reXjtsv4aahIomJUhsYNiVYUyXU5fS0vG1ti2JXBWX11yr4QapXlgQBLLaeVS1RkzokIT3KPe3djxQ5YPDlA2sNPT5MsMyKYLEFIYDj3OC3RnoWmAS+EUB8ggT96Pa9Pj687MYeN3WXambfuts2m8Rp3KUAV37QBbmlZfpo4QKDvkVHApyHpP4RTlMPyYAuZfNw6RTPFyuELHARQph0mqbqOjRRhpKKpTDYAy4VrUTJ2q5OVuMeSMGBLAgVLxrvhEluaVO7wKdM1js+Yb96wElGhGzClbtOjDDHDZPXdstLUBnYVikIboQpbjhlwyekxLjADoOk/8LQrl8fHECR4DiVVm/ZPorgFa6E6iVJt/1gHtoJZpxxrQ1rjA3hwbaK7Pf+V5Fg2KGbioPjgKhjhSaLi9LVyd3ZX/7qUfKC3dNuaUNwQ3/4aCuDNoe9z6NCxSTI8V5sH3z0znddhZUiVKd21tTCuB1rdB+DN5xQKRenPbeWsow6g2v6GlNlNSEyzn5UjlLlHyBmZ82uZWJjgfa+OGvVvAo4R4Fk0ZvV8NMUhchEtl5e4DltNeiObKVzU81lUQvZNlJBH0twChYWiXwD+GoZHhDjr8b5xSw==", keys[0].Key, "")
	assert.Equal(t, "Тест Юзер (rscz.ru)", keys[0].Email, "")
}
