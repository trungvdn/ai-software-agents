package dto

type StoryResponse struct {
	Stories []StoryItem `json:"stories"`
}

type StoryItem struct {
	Title  string `json:"title"`
	AsA    string `json:"as_a"`
	IWant  string `json:"i_want"`
	SoThat string `json:"so_that"`
}
