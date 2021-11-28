package web

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:required`
	Name string `json:"name" validate:required,min=0,max=100`
}
