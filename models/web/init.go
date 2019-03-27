package web

import (
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(new(Language), new(Tag), new(Videos), new(Category),new(PornStars))
}


// LanguageTBName 获取 Language 对应的表名称
func LanguageTBName() string {
	return "language"
}

// TagTBName 获取 Tag 对应的表名称
func TagTBName() string {
	return "tag"
}

func VideosTBName() string {
	return "videos"
}


func CategoryTBName()string{
	return "category"
}


func PornstarTBName()string{
	return "porn_stars"
}