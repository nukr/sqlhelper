package sqlhelper

import (
	"fmt"
	"io"
	"regexp"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(s *Scanner) *Parser {
	return &Parser{s: s}
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit
	return
}

func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) Parse(w io.Writer) error {
	for {
		tok, lit := p.scan()
		switch tok {
		case COMMENT_NAME:
			re := regexp.MustCompile(`--\ *name\ *:\ *(.*)`)
			variantName := string(re.FindSubmatch([]byte(lit))[1])
			fmt.Fprintf(w, "var %s = `\n", variantName)
		case STRING:
			fmt.Fprintf(w, "%s`\n\n", lit)
		case EOF:
			return nil
		}
	}
}
