package lyra2re_test

import (
	"fmt"
	"testing"

	"github.com/vertcoin/lyra2re"
)

func TestLyra2RE(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		out  string
	}{
		{
			name: "Basic",
			in:   []byte{0x34, 0x6e, 0x80, 0x88, 0x0e, 0xcc, 0x9e, 0x84, 0xce, 0x60, 0x22, 0xcf, 0x37, 0x56, 0xa1, 0xdf, 0x17, 0x56, 0x84, 0x0e, 0xf7, 0xea, 0x65, 0xc6, 0x44, 0xc9, 0x9f, 0x6d, 0x3d, 0xa3, 0x1f, 0x2b},
			out:  "90a631ac1f1f9f83fbd197bde457f20b5d73e7b8c02890f6c68dfcd9da3c6243",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			sum, err := lyra2re.Sum(test.in)
			if err != nil {
				t.Errorf("Error while summing: %v", err)
			}

			got := fmt.Sprintf("%x", sum)
			if got != test.out {
				t.Errorf("Expected %q", test.out)
				t.Errorf("Got %q", got)
			}
		})
	}
}
