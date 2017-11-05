package model

import "time"

type ResultItem struct {
	ID        int          `json:"id"`
	ResultID  int          `json:"resultID" db:"result_id"`
	TaskID    int          `json:"taskID" db:"task_id"`
	Output    string       `json:"output" db:"output"`
	Error     string       `json:"error"`
	Status    ResultStatus `json:"status"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at" `
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
}
