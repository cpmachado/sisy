package todotxt

// description definition
type Description struct {
	Text           string
	ProjectTags    []string
	ProjectContext []string
	CustomTags     []CustomTag
}

func (d *Description) String() string {
	return d.Text
}
