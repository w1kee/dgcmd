package dgcmd

import (
	"fmt"
	"strings"
	"unicode"
)

func validPrefix(s string) error {
	for i, r := range s {
		if r > unicode.MaxASCII {
			return fmt.Errorf("not unicode")
		}
		if r == ' ' && i != len([]rune(s))-1 {
			return fmt.Errorf("spaces in middle of prefix")
		}
	}
	return nil
}

func parseCommand(s string, p string) []string {
	if s[:len(p)] == p {
		if len(s) > len(p) {
			return strings.Split(s[len(p):], " ")
		}
	}
	return nil
}
