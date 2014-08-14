package extedit

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCreateContent(t *testing.T) {
	strRd := strings.NewReader("Line1\nLine2")

	c, err := contentFromReader(strRd, bufio.ScanLines)

	ok(t, err)
	equals(t, c.Length(), 2)
	equals(t, c.c[0], "Line1")
	equals(t, c.c[1], "Line2")
}

func TestCreateContentFromString(t *testing.T) {
	c, err := contentFromString("Line1\nLine2", bufio.ScanLines)

	ok(t, err)
	equals(t, c.Length(), 2)
	equals(t, c.c[0], "Line1")
	equals(t, c.c[1], "Line2")
}

func TestCreateContentFromFile(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	ok(t, err)
	_, err = f.Write([]byte("Line1\nLine2"))
	ok(t, err)

	c, err := contentFromFile(f.Name(), bufio.ScanLines)

	equals(t, c.Length(), 2)
	equals(t, c.c[0], "Line1")
	equals(t, c.c[1], "Line2")
}

func TestSplitByWords(t *testing.T) {
	c, err := contentFromString("Word1 Word2", bufio.ScanWords)

	ok(t, err)
	equals(t, c.Length(), 2)
	equals(t, c.c[0], "Word1")
	equals(t, c.c[1], "Word2")
}

func TestUseContentAsString(t *testing.T) {
	c, err := contentFromString("Line1\nLine2", bufio.ScanLines)

	ok(t, err)
	equals(t, fmt.Sprintf("x%sx", c), "xLine1\nLine2x")
}

func TestUseContentAsReader(t *testing.T) {
	c, err := contentFromString("Line1\nLine2", bufio.ScanLines)

	ok(t, err)
	byteCnt, err := ioutil.ReadAll(c)
	ok(t, err)
	equals(t, string(byteCnt), "Line1\nLine2")
}
