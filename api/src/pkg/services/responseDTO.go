package services

type NewCasesResponse struct {
	Year       int     `json:"year"`
	Month      int     `json:"month"`
	Location   string  `json:"location"`
	TotalCases float32 `json:"total_cases"`
	NewCases   float32 `json:"new_cases"`
}
