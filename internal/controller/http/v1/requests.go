package v1

type IdPathParams struct {
	ID string `uri:"id" binding:"required,uuid" format:"uuid"`
}
