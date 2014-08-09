// extedit provides functionality to open an editor with to let the user edit
// some content and get the changes the user made to process them as
// part of the user iterface of a command line programm.
package extedit

import (
  "io"
  "io/ioutil"
  "os"
  "os/exec"
	"bufio"
)

const default_editor = "vim"

// Session represents 
type Session struct {
	input Content
	result Content
	SplitFunc bufio.SplitFunc
}

// Session.Invoke starts a text-editor with the contents of content.
// After the user has closed the editor Invoke returns an
// io.Reader with the edited content.
func (s *Session) Invoke(content io.Reader) (Diff, error) {
	d := Diff{}

	input, err := contentFromReader(content, s.SplitFunc)
  if err != nil {
    return d, err
  }

	fileName, err := writeTmpFile(input)
  if err != nil {
    return d, err
  }

	cmd := editorCmd(fileName)
	err = cmd.Run()
  if err != nil {
    return d, err
  }

	result, err := contentFromFile(fileName, s.SplitFunc)
  if err != nil {
    return d, err
  }

	return NewDiff(input, result), nil
}

func NewSession() (*Session) {
	return &Session{SplitFunc: bufio.ScanLines}
}

// Invoke is a shortcut to the Invoke method of a default session.
func Invoke(content io.Reader) (Diff, error) {
	s := NewSession()
	return s.Invoke(content)
}

// writeTmpFile writes content to a temporary file and returns
// the path to the file
func writeTmpFile(content io.Reader) (string, error) {
  f, err := ioutil.TempFile("","")

  if err != nil {
    return "", err
  }

  io.Copy(f, content)
  f.Close()
	return f.Name(), nil
}

// editorCmd creates a os/exec.Cmd to open
// filename in an editor ready to be run()
func editorCmd(filename string) *exec.Cmd {
  editorPath := os.Getenv("EDITOR")
  if editorPath == "" {
    editorPath = default_editor
  }
  editor := exec.Command(editorPath, filename)

  editor.Stdin = os.Stdin
  editor.Stdout = os.Stdout
  editor.Stderr = os.Stderr

  return editor
}
