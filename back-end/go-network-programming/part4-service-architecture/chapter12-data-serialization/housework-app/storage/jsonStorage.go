package storage

import (
	"encoding/json"
	"housework-app/housework"
	"io"
)

type JsonStorage struct{}

func (s *JsonStorage) Load(r io.Reader) ([]*housework.Chore, error) {
	var chores []*housework.Chore
	return chores, json.NewDecoder(r).Decode(&chores)
}

func (s *JsonStorage) Flush(w io.Writer, chores []*housework.Chore) error {
	return json.NewEncoder(w).Encode(chores)
}
