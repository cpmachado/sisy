package todotxt

import "regexp"

// description type, parses tags and context and customtags lazily
type Description struct {
	Text    string
	parsed  bool
	project []string
	context []string
	special map[string]any
}

func (d *Description) String() string {
	return d.Text
}

func (d *Description) ProjectTags() []string {
	if !d.parsed {
		parseDescription(d)
	}

	return d.project
}

func (d *Description) ContextTags() []string {
	if !d.parsed {
		parseDescription(d)
	}

	return d.context
}

func parseDescription(d *Description) {
	if d.parsed {
		return
	}
	projectRE := regexp.MustCompile(`\+[^\s]+`)
	contextRE := regexp.MustCompile(`@[^\s]+`)

	for _, tag := range projectRE.FindAllString(d.Text, -1) {
		d.project = append(d.project, tag[1:])
	}

	for _, ctx := range contextRE.FindAllString(d.Text, -1) {
		d.context = append(d.context, ctx[1:])
	}

	d.parsed = true
}
