package api

import (
	"encoding/json"

	"gitlab.com/q-dev/q-client/rpc/shared"
)

type Sha3Args struct {
	Data string
}

func (args *Sha3Args) UnmarshalJSON(b []byte) (err error) {
	var obj []interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		return shared.NewDecodeParamError(err.Error())
	}

	if len(obj) < 1 {
		return shared.NewInsufficientParamsError(len(obj), 1)
	}

	argstr, ok := obj[0].(string)
	if !ok {
		return shared.NewInvalidTypeError("data", "is not a string")
	}
	args.Data = argstr
	return nil
}
