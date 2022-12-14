package post

type Post struct {
	id      int
	title   string
	content string
}

// Getter

func (p *Post) ID() int {
	return p.id
}

func (p *Post) Title() string {
	return p.title
}

func (p *Post) Content() string {
	return p.content
}

// Setter

func (p *Post) SetID(id int) {
	p.id = id
}

func (p *Post) SetTitle(title string) {
	p.title = title
}

func (p *Post) SetContent(content string) {
	p.content = content
}
