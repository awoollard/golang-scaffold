package model

// StatusResponse struct
type StatusResponse struct {
	Author []StatusItem `json:"my-application"`
}

// StatusItem struct
type StatusItem struct {
	Description string `json:"description"`
	Sha         string `json:"sha"`
	Version     string `json:"version"`
}

// MetaItem struct
type MetaItem struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}
