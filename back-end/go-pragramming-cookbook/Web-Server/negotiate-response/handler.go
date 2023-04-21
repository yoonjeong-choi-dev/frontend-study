package negotiate_response

import (
	"encoding/xml"
	"net/http"
)

type Payload struct {
	XMLName xml.Name `xml:"payload" json:"-"`
	Status  string   `xml:"status" json:"status"`
	Name    string   `xml:"name" json:"name"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	n := GetNegotiator(r)
	p := Payload{Name: "negotiator", Status: "Successful"}

	n.Respond(w, http.StatusOK, p)
}
