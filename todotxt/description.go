package todotxt

import "strings"

// description definition
type Description struct {
	Text           string
	ProjectTags    []string
	ProjectContext []string
	CustomTags     []CustomTag
}

func (d *Description) String() string {
	sb := &strings.Builder{}

	sb.WriteString(d.Text)

	for _, tag := range d.ProjectTags {
		sb.WriteString(" +")
		sb.WriteString(tag)
	}

	for _, ctx := range d.ProjectContext {
		sb.WriteString(" @")
		sb.WriteString(ctx)
	}

	for _, tag := range d.CustomTags {
		sb.WriteString(" ")
		sb.WriteString(tag.String())
	}

	return sb.String()
}
