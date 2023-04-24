package e2e_with_cucumber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"net/http/httptest"
)

var payloads []HandlerRequest
var responses []*httptest.ResponseRecorder

func aUserRequestGET() error {
	responses = make([]*httptest.ResponseRecorder, 0)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	Handler(w, r)
	responses = append(responses, w)
	return nil
}

func aUserRequestPOSTWithEmptyPayload() error {
	responses = make([]*httptest.ResponseRecorder, 0)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	Handler(w, r)
	responses = append(responses, w)
	return nil
}

func aUserRequestPOSTWithPayload(arg1 *godog.Table) error {
	payloads = make([]HandlerRequest, 0)
	responses = make([]*httptest.ResponseRecorder, 0)

	for _, row := range arg1.Rows {
		h := HandlerRequest{
			Name: row.Cells[0].Value,
		}
		payloads = append(payloads, h)
	}

	for _, p := range payloads {
		v, err := json.Marshal(p)
		if err != nil {
			return err
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer(v))
		Handler(w, r)
		responses = append(responses, w)
	}
	return nil
}

func theResponseBodyShouldBe(arg1 *godog.Table) error {
	for c, row := range arg1.Rows {
		b := bytes.Buffer{}
		b.ReadFrom(responses[c].Body)
		if got, want := b.String(), row.Cells[0].Value; got != want {
			return fmt.Errorf("got: %s, wnat: %s", got, want)
		}
	}
	return nil
}

func theResponseCodeShouldBe(arg1 int) error {
	for _, r := range responses {
		if got, want := r.Code, arg1; got != want {
			return fmt.Errorf("got: %d, wnat: %d", got, want)
		}
	}
	return nil
}

func PostGoodPayload(ctx *godog.ScenarioContext) {
	ctx.Step(`^a user request POST with payload:$`, aUserRequestPOSTWithPayload)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response body should be:$`, theResponseBodyShouldBe)
}

func PostEmptyPayload(ctx *godog.ScenarioContext) {
	ctx.Step(`^a user request POST with empty payload$`, aUserRequestPOSTWithEmptyPayload)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
}

func GetRequest(ctx *godog.ScenarioContext) {
	ctx.Step(`^a user request GET$`, aUserRequestGET)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
}
