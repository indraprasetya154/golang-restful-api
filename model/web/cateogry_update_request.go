package web

type CategoryUpdateRequest struct {
	Id   int    `validate:required`
	Name string `validate:required,min=0,max=100`
}
