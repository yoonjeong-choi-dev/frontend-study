package mongodb

import "context"

func MongoExample(host string) error {
	db, err := Setup(context.Background(), host)
	if err != nil {
		return err
	}
	defer db.cancel()

	data := []Data{
		{
			Name: "yoonjeong-choi-dev",
			Age:  31,
			Location: Location{
				City:    "Seoul",
				Country: "Korea",
			},
		},
		{
			Name: "yj",
			Age:  29,
			Location: Location{
				City:    "NewYork",
				Country: "USA",
			},
		},
	}

	if err := SaveManyData(db, data); err != nil {
		return err
	}

	if err := QueryByName(db, "yoonjeong-choi-dev"); err != nil {
		return err
	}
	if err := QueryByName(db, "yj"); err != nil {
		return err
	}

	// Clean up the database
	if err := DropByContext(db); err != nil {
		return err
	}

	return nil
}
