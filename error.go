package livy

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type errorMsg struct {
	Msg string `json:"msg"`
}

func extractErrorMessage(body []byte) error {
	var err errorMsg
	if err := json.Unmarshal(body, &err); err != nil {
		return errors.Wrap(err, "failed to decode response body")
	}

	return errors.New(err.Msg)
}
