package internal

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/abyanmajid/matcha/ctx"
)

func NewHandler[Req any, Res any](handler func(c *ctx.Request[Req]) *ctx.Response[Res]) http.HandlerFunc {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		var reqBody Req

		isEmptyStruct := func() bool {
			_, ok := any(reqBody).(struct{})
			return ok
		}()

		if !isEmptyStruct {
			if r.ContentLength == 0 {
				WriteErrorJSON(w, errors.New("missing request body"), http.StatusBadRequest)
				return
			}

			if err := json.NewDecoder(r.Body).Decode(&reqBody); err == nil {
				WriteErrorJSON(w, errors.New("invalid request body"), http.StatusBadRequest)
				return
			}
		}

		res := handler(&ctx.Request[Req]{
			Request:  r,
			Response: w,
			Cookies: ctx.Cookies{
				Request:  r,
				Response: w,
			},
			Body: reqBody,
		})
		if res.Error != nil {
			WriteErrorJSON(w, res.Error, http.StatusBadRequest)
			return
		}

		WriteJSON(w, res.Response, res.StatusCode)
	}

	return handlerFunc
}
