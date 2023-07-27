package main

import (
	"errors"
	"fmt"
	"housework-app/housework"
	"housework-app/storage"
	"os"
	"strconv"
	"strings"
)

type houseworkApp struct {
	storage  storage.HouseWorkStorage
	dataFile string
}

func (app *houseworkApp) load() ([]*housework.Chore, error) {
	if _, err := os.Stat(app.dataFile); os.IsNotExist(err) {
		return make([]*housework.Chore, 0), nil
	}

	file, err := os.Open(app.dataFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("cloasing data file: %v\n", err)
		}
	}()

	return app.storage.Load(file)
}

func (app *houseworkApp) flush(chores []*housework.Chore) error {
	file, err := os.Create(app.dataFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("cloasing data file: %v\n", err)
		}
	}()

	return app.storage.Flush(file, chores)
}

func (app *houseworkApp) printList() error {
	chores, err := app.load()
	if err != nil {
		return err
	}

	if len(chores) == 0 {
		fmt.Println("You are all caught up!")
		return nil
	}

	fmt.Println("#\t[X]\tDescription")
	for i, chore := range chores {
		complete := ""
		if chore.Complete {
			complete = "X"
		}
		fmt.Printf("%d\t[%s]\t%s\n", i+1, complete, chore.Description)
	}

	return nil
}

func (app *houseworkApp) add(s string) error {
	chores, err := app.load()
	if err != nil {
		return err
	}

	for _, chore := range strings.Split(s, ",") {
		if desc := strings.TrimSpace(chore); desc != "" {
			chores = append(chores, &housework.Chore{Description: desc})
		}
	}

	return app.flush(chores)
}

func (app *houseworkApp) complete(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	chores, err := app.load()
	if err != nil {
		return err
	}

	if i < 1 || i > len(chores) {
		return fmt.Errorf("chore %d not found", i)
	}

	chores[i-1].Complete = true
	return app.flush(chores)
}

func NewHouseworkApp(format, db string) (houseworkApp, error) {
	app := houseworkApp{}
	switch format {
	case "json":
		app.storage = &storage.JsonStorage{}
	case "gob":
		app.storage = &storage.GobStorage{}
	case "protobuf":
		app.storage = &storage.ProtobufStorage{}
	default:
		return app, errors.New("unsupported encoding format")
	}

	app.dataFile = db
	return app, nil
}
