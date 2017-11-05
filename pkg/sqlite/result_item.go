package sqlite

import (
	"github.com/jvikstedt/watchful/pkg/model"
)

func (s *sqlite) ResultItemCreate(resultItem *model.ResultItem) error {
	r, err := s.q.Exec(`INSERT INTO result_items (result_id, task_id, output, error, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, resultItem.ResultID, resultItem.TaskID, resultItem.Output, resultItem.Error, resultItem.Status)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}

	return s.ResultItemGetOne(int(id), resultItem)
}

func (s *sqlite) ResultItemGetOne(id int, resultItem *model.ResultItem) error {
	return s.q.Get(resultItem, "SELECT * FROM result_items WHERE id=$1", id)
}

func (s *sqlite) ResultItemUpdate(resultItem *model.ResultItem) error {
	_, err := s.q.Exec(`UPDATE result_items SET result_id = ?, task_id = ?, output = ?, error = ?, status = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, resultItem.ResultID, resultItem.TaskID, resultItem.Output, resultItem.Error, resultItem.Status, resultItem.ID)
	if err != nil {
		return err
	}

	return s.ResultItemGetOne(resultItem.ID, resultItem)
}

func (s *sqlite) ResultItemAllByResultID(resultID int) ([]*model.ResultItem, error) {
	resultItems := []*model.ResultItem{}
	err := s.q.Select(&resultItems, "SELECT * FROM result_items WHERE result_id = ?", resultID)
	return resultItems, err
}
