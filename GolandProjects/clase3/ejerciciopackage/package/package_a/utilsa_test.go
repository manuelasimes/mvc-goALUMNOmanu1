package package_a

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	assert.Equal(t, "Hello, world.", Hello())
}
