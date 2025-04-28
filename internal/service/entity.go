package service

// TaskRequest - структура, представляющая тело запроса
type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type TaskResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
}
