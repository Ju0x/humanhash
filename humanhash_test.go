package humanhash_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/Ju0x/humanhash"
)

func TestHumanHash(t *testing.T) {
	s := "Test"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	r, err := humanhash.Humanize(bs, 4, "-")

	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}
