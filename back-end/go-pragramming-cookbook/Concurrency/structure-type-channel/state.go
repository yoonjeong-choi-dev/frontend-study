package structure_type_channel

import "errors"

type op string

const (
	Add      op = "add"
	Subtract op = "sub"
	Multiply op = "mult"
	Divide   op = "div"
)

type OpsRequest struct {
	Operation op
	Value1    int64
	Value2    int64
}

type OpsResponse struct {
	Request *OpsRequest
	Result  int64
	Err     error
}

func Process(r *OpsRequest) *OpsResponse {
	res := OpsResponse{Request: r}

	switch r.Operation {
	case Add:
		res.Result = r.Value1 + r.Value2
	case Subtract:
		res.Result = r.Value1 - r.Value2
	case Multiply:
		res.Result = r.Value1 * r.Value2
	case Divide:
		if r.Value2 == 0 {
			res.Err = errors.New("divide by 0")
			break
		}
		res.Result = r.Value1 / r.Value2
	default:
		res.Err = errors.New("unsupported operation")
	}

	return &res
}
