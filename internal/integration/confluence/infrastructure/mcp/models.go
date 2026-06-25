package mcp

type CreatePageRequest struct {
	SpaceKey string
	ParentID string
	Title    string
	Content  string
}

type CreatePageResponse struct {
	ID      string
	URL     string
	Version int
}
