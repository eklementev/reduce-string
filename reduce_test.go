package reducestring_test

import (
	"testing"

	"github.com/eklementev/reducestring"
	"github.com/stretchr/testify/assert"
)

func TestSanity(t *testing.T) {
	cases := map[string]string{
		"":                 "",
		"aaabbbccccc":      "a3b3c5",
		"aaabbbcccccaaaaa": "a8b3c5",
		"zzzzcccUUUuu":     "U3c3u2z4",
		"ЯЯЯБББдд":         "Б3Я3д2",
	}

	for source, expected := range cases {
		actual := reducestring.Reduce(source)
		assert.Equal(t, expected, actual)
	}
}
