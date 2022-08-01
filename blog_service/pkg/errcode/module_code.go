package errcode

var (
	GetListFail   = NewError(200001, "获取标签列表失败")
	CreateTagFail = NewError(200002, "创建标签失败")
	UpdateTagFail = NewError(200003, "更新标签失败")
	DeleteTagFail = NewError(200004, "删除标签失败")
	CountTagFail  = NewError(200005, "统计标签失败")

	ErrGetArticle    = NewError(300001, "获取单个文章失败")
	ErrGetArticles   = NewError(300002, "获取文章列表失败")
	ErrCreateArticle = NewError(300003, "创建文章失败")
	ErrUpdateArticle = NewError(300004, "更新文章失败")
	ErrDeleteArticle = NewError(30005, "删除文章失败")
)
