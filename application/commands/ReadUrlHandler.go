package commands

import (
	"sur/application/messages"
	"sur/domain/entities"
	"sur/domain/ports"
)

type ReadUrlHandler struct {
	repository ports.IUrlRepository
}

// ****** constructor
func NewReadUrlHandler(r ports.IUrlRepository) *ReadUrlHandler {
	return &ReadUrlHandler{
		repository: r,
	}
}

// ****** Functions
func (h *ReadUrlHandler) Handler(request messages.ReadUrlRequest) (*messages.ReadUrlResponse, error) {
	tenant := entities.Tentant{}
	code := request.Code
	url, _ := h.repository.Read(tenant, code)
	// 2. Create url code
	return &messages.ReadUrlResponse{
		Url: url.Url,
	}, nil
}
