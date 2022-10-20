package endpoints

import (
	"encoding/json"
	"net/http"
	"sur/application/commands"
	"sur/application/messages"
)

type CreateUrl struct {
	command commands.CreateUrlHandler
}

func (c *CreateUrl) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var request messages.CreateUrlRequest
	json.NewDecoder(r.Body).Decode((&request))
	response, err := c.command.Handler(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func NewCreateUrl(h commands.CreateUrlHandler) *CreateUrl {
	return &CreateUrl{
		command: h,
	}
}
