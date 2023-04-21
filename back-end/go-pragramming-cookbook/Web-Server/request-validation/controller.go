package request_validation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Controller 검증 함수를 가지고 있는 컨트롤러
// => 컨트롤러 초기화 시, 유연하게 검증 함수 등록 가능
type Controller struct {
	ValidatePayload func(p *Payload) error
}

func New() *Controller {
	return &Controller{
		ValidatePayload: ValidatePayload,
	}
}

func (c *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer func() { _ = r.Body.Close() }()

	var p Payload
	if err := decoder.Decode(&p); err != nil {
		log.Printf("Body Parsing Error: %s\n", err.Error())
		http.Error(w, "invalid data - must have name and age", http.StatusBadRequest)
		return
	}

	if err := c.ValidatePayload(&p); err != nil {
		switch err.(type) {
		case ValidationError:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		default:
			log.Printf("Validation Error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hi~ %s! Your age is %d", p.Name, p.Age)))
}
