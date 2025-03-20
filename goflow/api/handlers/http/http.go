package internal_http

import (
	// other imports
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateBody(request *http.Request, decode struct {}) error {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}
	_ = json.Unmarshal([]byte(body),&decode)
	validateBody := validator.New()
	err = validateBody.Struct(decode)
	return err
}