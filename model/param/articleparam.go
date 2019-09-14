package param

type ArticleParam struct {
	Id      uint   `form:"id"`
	Title   string `form:"title" binding:"required"`
	Alias   string `form:"alias"`
	Content string `form:"content" binding:"required"`
}
