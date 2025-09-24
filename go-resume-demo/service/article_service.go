package service

import "github.com/Cyte-091/go-resume-demo/errno"

func CreateArticle(title string) error {
	if title == "" {
		return errno.New(errno.ParamErr, "标题不能为空")
	}
	// 正常逻辑
	return nil
}
