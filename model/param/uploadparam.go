package param

type UploadParam struct {
	Filename string `form:"Filename"  binding:"required"`
	Created  string `form:"Created" binding:"required"`
}
