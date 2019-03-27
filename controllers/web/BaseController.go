package web

import (
	"fmt"
	m "github.com/MobileCPX/PreBobsTube/models/web"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//
type BaseController struct {
	// 内容站语言
	Language string
	// 登录状态
	UserSession string
	// 为每一个新的请求建一个连接，不用每次查数据都重新建立新数据库链接
	Orm orm.Ormer
	// LanguageAdaptation 内容站语言适配
	LanguageAdaption *m.Language

	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Language = beego.AppConfig.String("language")
	// 初始化数据库链接
	c.initORM()
	// 检查用户是否登录
	c.checkLogin()
	// 语言适配
	language := new(m.Language)
	c.LanguageAdaption, _ = language.LanguageAdaption(c.Language, c.Orm)
}

// 初始化数据库链接
func (c *BaseController) initORM() {
	c.Orm = orm.NewOrm()
}

// 检查用户是否登录
func (c *BaseController) checkLogin() (sessionStatus string) {
	sessionName := beego.AppConfig.String("session_name")
	c.Data["LoginStatus"] = true
	if sessionName == "" {
		sessionName = "user_id"
	}
	if c.Ctx.Input.Session(sessionName) == nil {
		c.UserSession = ""

	} else {
		c.Data["LoginStatus"] = true
		c.UserSession = c.Ctx.Input.Session(sessionName).(string)
	}

	return c.UserSession
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "temp/layout/layout_page.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	}
	c.Layout = layout
	c.TplName = tplName
}

func (c *BaseController) setRandomTags() {
	tags, _ := m.RandomTags(c.Orm, 30)
	c.Data["RandomTag"] = &tags
}

// 获取当前页码
func (c *BaseController) currentPageNum()int{
	currentPage, err := c.GetInt("p", 1)
	if err != nil {
		currentPage = 1
	}
	return currentPage
}

func (c *BaseController) leftSidebarSort(sortType string) {
	var sidebarDataList []map[string]interface{}
	topCategories, _ := m.GetTopCategory(c.Orm, "10")
	topPornstarts, _ := m.GetTopPornstars(c.Orm, "10")
	topTags, _ := m.GetTopTags(c.Orm, "10")

	categorys := make(map[string]interface{})
	categorys["type"] = "Top Categories"
	categorys["type_id"] = "categories"
	categorys["data"] = topCategories

	tags := make(map[string]interface{})
	tags["type"] = "Top Tags"
	tags["type_id"] = "tags"
	tags["data"] = topTags

	pornstars := make(map[string]interface{})
	pornstars["type"] = "Top Pornstars"
	pornstars["data"] = topPornstarts
	pornstars["type_id"] = "pornstars"

	sidebarDataList = leftSidebarSort(categorys, tags, pornstars, sortType)

	c.Data["LeftSidebar"] = sidebarDataList
}

func leftSidebarSort(categorys, tags, pornstars map[string]interface{}, sortType string) []map[string]interface{} {
	var sidebarDataList []map[string]interface{}
	switch sortType {
	case "tag":
		sidebarDataList = append(sidebarDataList, tags, categorys, pornstars)
	case "pornstar":
		sidebarDataList = append(sidebarDataList, pornstars, categorys, tags)
	default:
		sidebarDataList = append(sidebarDataList, categorys, pornstars, tags)
	}
	return sidebarDataList
}

func (c *BaseController) setRandomCatagory() {
	categorys, _ := m.TopCategories(c.Orm, 10)
	c.Data["TopCategories"] = categorys
}

func (c *BaseController) redirect(URL string) {
	c.Redirect(URL, 302)
	c.StopRun()
}

// currentPage 当前页数 totalPage 合计页数,
// redirectURL  跳转到指定页面的url 格式为  '/a/b?p=%d'   %d 格式化页码
func (c *BaseController) setPagination(currentPage, totalPage int, redirectURL string) string {
	pageHtml := ""

	if currentPage != 1 {
		pageHtml = fmt.Sprintf(`<ul>
					 <li class="prev"><a href="%s">Back</a></li>
					 <li class="first"><a href="%s">First</a></li>
					 <li class="jump"><a href="%s">...</a></li>
	`, pageURLFormat(redirectURL, currentPage-1), pageURLFormat(redirectURL, 1), "#")
	} else {
		pageHtml = `	<ul>
					 <li class="prev"><span>Back</span></li>
					 <li class="first"><span>First</span></li>
					 <li class="jump"><span>...</span></li>
	`
	}

	if currentPage+4 <= totalPage && currentPage-4 >= 1 {
		for i := currentPage - 4; i <= currentPage+4; i++ {
			if i == currentPage {
				pageHtml += fmt.Sprintf(`<li class="page-current"><span>%d</span></li>`, i)
			} else {
				pageHtml += fmt.Sprintf(`<li class="page"><a href="%s">%d</a></li>`, pageURLFormat(redirectURL, i), i)
			}
		}
	} else if currentPage-4 < 1 {
		endPage := 9
		if totalPage <= 9 {
			endPage = totalPage
		}
		for i := 1; i <= endPage; i++ {
			if i == currentPage {
				pageHtml += fmt.Sprintf(`<li class="page-current"><span>%d</span></li>`, i)
			} else {
				pageHtml += fmt.Sprintf(`<li class="page"><a href="%s">%d</a></li>`, pageURLFormat(redirectURL, i), i)
			}
		}
	} else if currentPage+4 >= totalPage {
		startPage := totalPage - 9
		if totalPage-9 <= 0 {
			startPage = 1
		}
		for i := startPage; i <= totalPage; i++ {
			if i == currentPage {
				pageHtml += fmt.Sprintf(`<li class="page-current"><span>%d</span></li>`, i)
			} else {
				pageHtml += fmt.Sprintf(`<li class="page"><a href="%s">%d</a></li>`, pageURLFormat(redirectURL, i), i)
			}
		}
	}

	pageHtml += fmt.Sprintf(`
 				<li class="jump"><a href="%s">...</a></li>
				<li class="last"><a href="%s">Last</a></li>
				<li class="next"><a href="%s">Next</a></li>
	`, "#", pageURLFormat(redirectURL, totalPage), pageURLFormat(redirectURL, currentPage+1))

	c.Data["Pagination"] = pageHtml

	return pageHtml
}

// 格式化跳转页面URL
// 例如 URL =  "/new?p=%d"   pageNum = 2   返回数据为 "/new?p=2"
func pageURLFormat(URL string, pageNum int) string {
	return fmt.Sprintf(URL, pageNum)
}
