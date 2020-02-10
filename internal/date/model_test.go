package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestShouldMarshalJSON(t *testing.T) {
	d := Date(time.Date(2049,6,3,11,22,33,44, time.UTC))
	bytes, err := d.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"2049-06-03\"", string(bytes))
}

func TestShouldUnmarshalJSON(t *testing.T) {
	d := Date{}
	err := d.UnmarshalJSON([]byte("\"2049-06-03\""))
	assert.Nil(t, err)
	assert.Equal(t, "2049-06-03", d.Format(Format))
}
