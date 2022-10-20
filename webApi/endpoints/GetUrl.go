package endpoints

import (
	"encoding/json"
	"net/http"
	"sur/application/commands"
	"sur/application/messages"

	"github.com/gorilla/mux"
)

type GetUrl struct {
	command commands.ReadUrlHandler
}

// ****** constructor
func NewGetUrl(h commands.ReadUrlHandler) *GetUrl {
	return &GetUrl{
		command: h,
	}
}

// ****** functions
func (g *GetUrl) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	code := mux.Vars(r)["id"]
	request := messages.ReadUrlRequest{
		Code: code,
	}
	response, err := g.command.Handler(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	http.Redirect(w, r, response.Url, http.StatusSeeOther)

}
