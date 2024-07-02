package v1

import "github.com/Anwarjondev/blog-website-clone/storage"

type handlerV1 struct {
	strg storage.StorageI
}
type HandlerV1 struct {
	Strg storage.StorageI
}

func New(h *HandlerV1) *handlerV1 {
	return &handlerV1{
		strg: h.Strg,
	}

}
