package services

type NewCasesResponse struct {
	Year        int    `json:"year"`
	Month       int    `json:"month"`
	Location    string `json:"location"`
	Total_Cases int    `json:"total_cases"`
	New_Cases   int    `json:"new_cases"`
}
