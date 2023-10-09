package humanhash_test

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"testing"

	"github.com/Ju0x/humanhash"
)

func TestHumanHash(t *testing.T) {
	generateHumanHash(t)
	generateHumanHashUUID(t)
}

func generateHumanHash(t *testing.T) {
	s := "Test"

	humanhash.Wordlist("example_wordlist.txt")

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	r, err := humanhash.Humanize(bs, 4, "-")

	if err != nil {
		t.Error(err)
	}

	if r != "autumn-papa-muppet-aspen" {
		t.Error(errors.New("hash not matching"))
	}
}

func generateHumanHashUUID(t *testing.T) {
	hash, uuid, err := humanhash.UUID()

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Humanhash: %s UUID: %s\n", hash, uuid)
}
