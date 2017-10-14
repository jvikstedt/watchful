package model

import (
	"encoding/json"
	"time"
)

type Job struct {
	ID         int                    `json:"id"`
	ExecutorID string                 `json:"executorID" db:"executor_id"`
	Takes      map[string]interface{} `json:"takes"`
	TakesRaw   []byte                 `json:"-" db:"takes_raw"`
	CreatedAt  time.Time              `json:"createdAt" db:"created_at" `
	UpdatedAt  time.Time              `json:"updatedAt" db:"updated_at"`
}

func (j *Job) Serialize() error {
	bytes, err := json.Marshal(j.Takes)
	if err != nil {
		return err
	}
	j.TakesRaw = bytes
	return nil
}

func (j *Job) Deserialize() error {
	takes := map[string]interface{}{}
	err := json.Unmarshal(j.TakesRaw, &takes)
	if err != nil {
		return err
	}
	j.Takes = takes
	return nil
}
