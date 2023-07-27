package storage

import (
	"encoding/gob"
	"housework-app/housework"
	"io"
)

type GobStorage struct{}

func (s *GobStorage) Load(r io.Reader) ([]*housework.Chore, error) {
	var chores []*housework.Chore
	return chores, gob.NewDecoder(r).Decode(&chores)
}

func (s *GobStorage) Flush(w io.Writer, chores []*housework.Chore) error {
	return gob.NewEncoder(w).Encode(chores)
}
