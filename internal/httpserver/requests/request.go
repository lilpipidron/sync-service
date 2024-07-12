package requests

import (
	"github.com/go-chi/render"
	"net/http"
)

func Decode(w http.ResponseWriter, r *http.Request, request *interface{}) error {
	err := render.DecodeJSON(r.Body, &request)
	return err
}
