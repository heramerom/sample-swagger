package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Scanner struct {
	reader *bufio.Reader
	f      string
	l      int
}

const eof = -1

func newScanner(s, f string, l int) *Scanner {
	return &Scanner{
		reader: bufio.NewReader(bytes.NewBufferString(s)),
		f:      f,
		l:      l,
	}
}

func (s *Scanner) next() int {
	ch, err := s.reader.ReadByte()
	if err != nil {
		return eof
	}
	return int(ch)
}

func (s *Scanner) isSep(ch int, sep ...byte) bool {
	for _, b := range sep {
		if ch == int(b) {
			return true
		}
	}
	return false
}

func (s *Scanner) skipWriteSpace() int {
	ch := s.next()
	for ; ch == ' ' || ch == '\t'; ch = s.next() {
	}
	return ch
}

func (s *Scanner) nextString(sep ...byte) string {
	buf := bytes.NewBuffer(nil)
	var quota byte
	ch := s.skipWriteSpace()
	if ch == '"' || ch == '\'' {
		quota = byte(ch)
		ch = s.skipWriteSpace()
	}
	for {
		if ch == eof {
			if quota != 0x00 {
				fmt.Println("syntax error: quotation not closed", s.f, s.l)
				os.Exit(1)
			}
			break
		}
		if byte(ch) == quota {
			break
		}
		if quota != 0x00 {
			goto Next
		}
		if !s.isSep(ch, sep...) {
			goto Next
		}
		break
	Next:
		buf.WriteByte(byte(ch))
		ch = s.next()
	}
	if !s.isSep(ch, sep...) {
		ch = s.skipWriteSpace()
		if !s.isSep(ch, sep...) && ch != eof {
			fmt.Println("syntax error:", s.f, s.l)
			os.Exit(1)
		}
	}
	return strings.TrimSpace(buf.String())
}
