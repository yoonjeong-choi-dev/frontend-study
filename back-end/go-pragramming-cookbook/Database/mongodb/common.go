package mongodb

type Location struct {
	City    string `bson:"city"`
	Country string `bson:"country"`
}

type Data struct {
	Name     string   `bson:"name"`
	Age      int      `bson:"age"`
	Location Location `bson:"location"`
}

const (
	Database   = "gocookbook"
	Collection = "example"
)
