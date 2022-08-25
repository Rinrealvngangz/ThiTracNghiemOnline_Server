package presenter

import "time"

type StudentResponse struct {
	Id        string
	Content   string
	CreatedAt time.Time
	CreatedBy string
}

type StudentRequest struct {
	Content string
}
