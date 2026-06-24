package dto

type EpicResponse struct {
	Epics []EpicItem `json:"epics"`
}

type EpicItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
