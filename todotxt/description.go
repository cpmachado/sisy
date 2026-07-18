package todotxt

// description definition
type Description struct {
	Text           string
	ProjectTags    []string
	ProjectContext []string
	CustomTags     []CustomTag
}
