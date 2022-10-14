package toolkit

import (
	"testing"
)

func TestTools_RandomStrings(t *testing.T) {

	var Tools Tools

	src := Tools.RandomString(10)

	if len(src) != 10 {
		t.Error("Wrong")
	}

}
