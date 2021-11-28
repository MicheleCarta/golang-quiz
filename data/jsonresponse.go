package data

type JsonResponse struct {
	Type    string   `json:"type"`
	Data    []Player `json:"data"`
	Message string   `json:"message"`
}
