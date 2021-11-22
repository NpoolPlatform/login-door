package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation(t *testing.T) {
	resp, err := GetLocationByIP("218.77.129.195")
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}
}
