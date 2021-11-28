package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:required,min=0,max=100`
}
