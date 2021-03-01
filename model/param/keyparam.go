package param

type KeyParam struct {
	Ip  string `form:"Ip"  binding:"required"`
	Key string `form:"Key" binding:"required"`
}
