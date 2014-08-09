package extedit

type Diff struct {
	Content Content
	Differences []int
}

func (d Diff) Lines() []string {
	return d.Content.c
}

func (d Diff) Line(i int) string {
	return d.Content.c[i]
}

func NewDiff(a, b Content) Diff {
	d := Diff{ Content: b }

	for i, line := range b.c {
		if len(a.c) <= i || line != a.c[i]  {
			d.Differences = append(d.Differences, i)
		}
	}

	return d
}
