package domain

type Page struct {
	ID       string
	Title    string
	Content  string
	URL      string
	ParentID string
	SpaceKey string
	Labels   []string
	Version  int
}
