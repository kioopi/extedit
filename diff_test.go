package extedit

import (
    "testing"
    "bufio"
)

func TestNoDiff(t *testing.T) {
  c1, err := contentFromString("Line1\nLine2", bufio.ScanLines)
  ok(t, err)

  c2, err := contentFromString("Line1\nLine2", bufio.ScanLines)
  ok(t, err)

	d := NewDiff(c1, c2)

  equals(t, len(d.Differences), 0)
}

func TestSimpleDiff(t *testing.T) {
  c1, err := contentFromString("Line1\nLine2", bufio.ScanLines)
  ok(t, err)

  c2, err := contentFromString("Line1\nXLine2", bufio.ScanLines)
  ok(t, err)

	d := NewDiff(c1, c2)

  equals(t, len(d.Differences), 1)

	equals(t, d.Line(d.Differences[0]), "XLine2")
}

func TestDiffExtraLines(t *testing.T) {
  c1, err := contentFromString("Line1\nLine2", bufio.ScanLines)
  ok(t, err)

  c2, err := contentFromString("Line1\nLine2\nLine3", bufio.ScanLines)
  ok(t, err)

	d := NewDiff(c1, c2)

  equals(t, len(d.Differences), 1)
	equals(t, d.Line(d.Differences[0]), "Line3")
}
