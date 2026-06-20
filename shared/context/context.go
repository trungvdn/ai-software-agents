package context

type PromptContext struct {
	Content string
}

type KnowledgeContext struct {
	Reflections    []string
	HistoricalBugs []string
}
