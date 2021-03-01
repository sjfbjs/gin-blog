package param

type KeyParam struct {
	Ip  string `form:"ip"  binding:"required"`
	Key string `form:"key" binding:"required"`
}
