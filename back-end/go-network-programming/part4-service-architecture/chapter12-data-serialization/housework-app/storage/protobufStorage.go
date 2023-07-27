package storage

import (
	"google.golang.org/protobuf/proto"
	"housework-app/housework"
	ph "housework-app/housework/v1"
	"io"
	"io/ioutil"
)

type ProtobufStorage struct{}

func (s *ProtobufStorage) Load(r io.Reader) ([]*housework.Chore, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var protos ph.Chores
	err = proto.Unmarshal(b, &protos)
	if err != nil {
		return nil, err
	}

	chores := protos.Chores

	ret := make([]*housework.Chore, len(chores))
	for i, chore := range chores {
		ret[i] = &housework.Chore{
			Complete:    chore.Complete,
			Description: chore.Description,
		}
	}

	return ret, nil
}

func (s *ProtobufStorage) Flush(w io.Writer, chores []*housework.Chore) error {

	phChores := make([]*ph.Chore, len(chores))
	for i, chore := range chores {
		phChores[i] = &ph.Chore{
			Complete:    chore.Complete,
			Description: chore.Description,
		}
	}

	b, err := proto.Marshal(&ph.Chores{Chores: phChores})
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}
