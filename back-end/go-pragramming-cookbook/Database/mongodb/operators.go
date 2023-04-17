package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func convertToInterfaceSlice(data []Data) []interface{} {
	ret := make([]interface{}, len(data))
	for i := range data {
		ret[i] = data[i]
	}
	return ret
}

func SaveManyData(db *DB, data []Data) error {
	toInsert := convertToInterfaceSlice(data)

	collection := db.client.Database(Database).Collection(Collection)
	if _, err := collection.InsertMany(db.ctx, toInsert); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func QueryByName(db *DB, name string) error {
	collection := db.client.Database(Database).Collection(Collection)

	var data Data
	if err := collection.FindOne(db.ctx, bson.M{"name": name}).Decode(&data); err != nil {
		return nil
	}

	fmt.Printf("Result by name '%s': %v\n", name, data)
	return nil
}

func DropByContext(db *DB) error {
	collection := db.client.Database(Database).Collection(Collection)
	return collection.Drop(db.ctx)
}
