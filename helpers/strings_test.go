package helpers

import (
	"testing"

	"github.com/facebookgo/ensure"
)

func TestUnquote(t *testing.T) {
	ensure.DeepEqual(t, UnquoteIfNeeded("abcdef", '-'), "abcdef")
	ensure.DeepEqual(t, UnquoteIfNeeded("'123456'", '\''), "123456")
	ensure.DeepEqual(t, UnquoteIfNeeded("<123456<", '<'), "123456")
	ensure.DeepEqual(t, UnquoteIfNeeded("<123456>", '<'), "<123456>")
}
