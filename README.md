# extedit

Open an external editor as part of the user interface of a command-line programm. Think `git rebase -i`.


```golang
import (
    "strings"
    "github.com/kioopi/extedit"
)

func main() {
    input := strings.NewReader("Line 1\nLine 2")

    diff := extedit.Invoke(input)

    # diff.Lines() []string contains the edited input.
    # diff.Differences []int  contains indexes of changed lines.
}
```
