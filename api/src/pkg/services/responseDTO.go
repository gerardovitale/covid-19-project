package services

type NewCasesResponse struct {
	Year       int    `json:"year"`
	Month      int    `json:"month"`
	Location   string `json:"location"`
	TotalCases int    `json:"total_cases"`
	NewCases   int    `json:"new_cases"`
}
