package web

type CategoryCreateRequest struct {
	Name string `validate:required,min=0,max=100`
}
