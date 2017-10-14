package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jvikstedt/watchful/model"
)

func (h handler) jobCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	job := &model.Job{}

	if err := decoder.Decode(job); h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}

	job.Serialize()

	newJob, err := h.storage.JobCreate(job.ExecutorID, job.TakesRaw)
	if h.checkErr(err, w, http.StatusUnprocessableEntity) {
		return
	}
	newJob.Deserialize()

	json.NewEncoder(w).Encode(newJob)
}
