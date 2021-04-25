package dgcmd

import (
	"reflect"
	"testing"
)

func TestGetCommandName(t *testing.T) {
	tests := []struct {
		s   string
		p   string
		naa []string
	}{
		{"!test", "!", []string{"test"}},
		{"! test", "! ", []string{"test"}},
		{"pls help with something", "pls ", []string{"help", "with", "something"}},
	}

	for i, test := range tests {
		naa := parseCommand(test.s, test.p)
		if !reflect.DeepEqual(naa, test.naa) {
			t.Errorf("%d: args are %#v, but should be %#v", i, naa, test.naa)
		}
	}
}
