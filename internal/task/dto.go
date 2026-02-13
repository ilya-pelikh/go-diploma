package task

type Task struct {
	Id      int64
	Date    string
	Title   string
	Comment string
	Repeat  string
}

type AddTaskRequestDTO struct {
	Date    string `json:"date" validate:"omitempty,datetime=20060102"`
	Title   string `json:"title" validate:"required"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}
type UpdateTaskRequestDTO struct {
	Id      int64  `json:"id,string" validate:"required"`
	Date    string `json:"date" validate:"omitempty,datetime=20060102"`
	Title   string `json:"title" validate:"required"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type TaskResponseDTO struct {
	Id      int64  `json:"id,string"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

type AddTaskResponseDTO struct {
	Id      int64 `json:"id,string"`
	date    string
	title   string
	comment string
	repeat  string
}
