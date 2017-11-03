package sqlhelper

import (
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	cases := []struct {
		s   string
		tok Token
		lit string
	}{
		{
			s:   ``,
			tok: EOF,
			lit: "",
		},
		{
			s: `
`,
			tok: WS,
			lit: "\n",
		},
		{
			s:   `-- name: ggg`,
			tok: COMMENT_NAME,
			lit: "-- name: ggg",
		},
		{
			s:   `-- ggg`,
			tok: COMMENT,
			lit: "-- ggg",
		},
		{
			s:   `ZZZ`,
			tok: STRING,
			lit: "ZZZ",
		},
	}
	for _, c := range cases {
		s := NewScanner(strings.NewReader(c.s))
		gotToken, gotLit := s.Scan()
		if gotToken != c.tok || gotLit != c.lit {
			t.Errorf(`
			want tok: %s
			got tok: %s
			want lit: %s
			got lit: %s
			`, c.tok.String(), gotToken.String(), c.lit, gotLit)
		}
	}
}
