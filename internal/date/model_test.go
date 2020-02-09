package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestDate_MarshalJSON(t *testing.T) {
	d := Date(time.Date(2049,6,3,11,22,33,44, time.UTC))
	bytes, err := d.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"2049-06-03\"", string(bytes))
}
