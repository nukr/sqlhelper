package sqlhelper

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
)

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS
	COMMENT
	COMMENT_NAME
	STRING
)

func (t Token) String() string {
	switch t {
	case 0:
		return "ILLEGAL"
	case 1:
		return "EOF"
	case 2:
		return "WS"
	case 3:
		return "IDENT"
	case 4:
		return "COMMENT"
	case 5:
		return "COMMENT_NAME"
	case 6:
		return "STRING"
	default:
		return "UNKNOWN"
	}
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return rune(0)
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }
func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if ch == '-' {
		next, _ := s.r.Peek(1)
		if rune(next[0]) == '-' {
			s.unread()
			return s.scanComment()
		} else {
			s.unread()
		}
	} else if ch != rune(0) {
		s.unread()
		return s.scanString()
	}
	switch ch {
	case rune(0):
		return EOF, ""
	}
	return ILLEGAL, string(ch)
}

func (s *Scanner) scanComment() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == rune(0) || ch == '\n' {
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	re := regexp.MustCompile(`--\ *name:\ *(.*)`)
	if re.Match(buf.Bytes()) {
		return COMMENT_NAME, buf.String()
	}
	return COMMENT, buf.String()
}

func (s *Scanner) scanString() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if ch == rune(0) {
			break
		} else if ch == '-' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}
	return STRING, buf.String()
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if ch := s.read(); ch == rune(0) {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
