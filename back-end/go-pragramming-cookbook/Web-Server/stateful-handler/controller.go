package stateful_handler

type Controller struct {
	storage Storage
}

func NewController(storage Storage) *Controller {
	return &Controller{storage: storage}
}

type Payload struct {
	Value string `json:"value"`
}
