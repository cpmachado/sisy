package todotxt

import "regexp"

// description type, parses tags and context and customtags lazily
type Description struct {
	Text    string
	parsed  bool
	project []string
	context []string
	special map[string]string
}

func (d *Description) Init(text string) {
	d.Text = text
	d.parsed = false
	d.project = nil
	d.context = nil
	d.special = nil
}

func (d *Description) String() string {
	if d == nil {
		return ""
	}
	return d.Text
}

func (d *Description) ProjectTags() []string {
	if d == nil {
		return nil
	}
	d.parseDescription()

	return d.project
}

func (d *Description) ContextTags() []string {
	if d == nil {
		return nil
	}
	d.parseDescription()

	return d.context
}

func (d *Description) Special() map[string]string {
	if d == nil {
		return nil
	}
	d.parseDescription()
	return d.special
}

func (d *Description) parseDescription() {
	if d == nil || d.parsed {
		return
	}
	projectRE := regexp.MustCompile(`\+[^\s]+`)
	contextRE := regexp.MustCompile(`@[^\s]+`)
	specialRE := regexp.MustCompile(`(\w+):(\S+)`)

	for _, tag := range projectRE.FindAllString(d.Text, -1) {
		d.project = append(d.project, tag[1:])
	}

	for _, ctx := range contextRE.FindAllString(d.Text, -1) {
		d.context = append(d.context, ctx[1:])
	}

	if d.special == nil {
		d.special = make(map[string]string)
	}

	for _, kv := range specialRE.FindAllStringSubmatch(d.Text, -1) {
		d.special[kv[1]] = kv[2]
	}

	d.parsed = true
}
