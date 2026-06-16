package bugfix

import "fmt"

type ReadFileTool interface {
	Read(path string) (string, error)
}

type SearchCodeTool interface {
	Search(keyword string) ([]string, error)
}

type Agent struct {
	readFile ReadFileTool
	search   SearchCodeTool
}

func New(r ReadFileTool, s SearchCodeTool) *Agent {
	return &Agent{
		readFile: r,
		search:   s,
	}
}

func (a *Agent) Run(query string) (string, error) {
	results, err := a.search.Search(query)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"query=%s\nmatched_files=%v",
		query,
		results,
	), nil
}
