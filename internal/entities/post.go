package entities

type Post struct {
	Title string
	Body  string
}

func (p *Post) GetTitle() string {
	return p.Title
}

func (p *Post) GetBody() string {
	return p.Body
}

func (p *Post) SetTitle(title string) {
	p.Title = title
}

func (p *Post) SetBody(body string) {
	p.Body = body
}
