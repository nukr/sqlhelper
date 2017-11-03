package sqlhelper

import (
	"bytes"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{
			s: `
-- name: gg
SELECT * FROM products
`,
			want: "var gg = `\nSELECT * FROM products\n`\n\n",
		},
	}

	for _, c := range cases {
		parser := NewParser(NewScanner(strings.NewReader(c.s)))
		s := bytes.NewBufferString("")
		parser.Parse(s)
		if c.want != s.String() {
			t.Errorf(`
			want: %s
			got: %s
			`, c.want, s.String())
		}
	}
}
