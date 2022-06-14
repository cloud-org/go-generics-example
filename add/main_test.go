package add

import (
	"testing"
)

func Test_add(t *testing.T) {
	t.Log(add(1, 2))
	t.Log(add("foo", "bar"))
}
