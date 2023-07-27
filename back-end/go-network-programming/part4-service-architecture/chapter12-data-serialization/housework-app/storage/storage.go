package storage

import (
	"housework-app/housework"
	"io"
)

type HouseWorkStorage interface {
	Load(r io.Reader) ([]*housework.Chore, error)
	Flush(w io.Writer, chores []*housework.Chore) error
}
