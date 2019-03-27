package routers

import (
	"github.com/MobileCPX/PreBobsTube/controllers/web"
	"github.com/astaxie/beego"
)

func init() {
	// 首页
	beego.Router("/", &web.HomeController{})

	// 其他页面
	beego.Router("/latest-updates", &web.OtherPageController{}, "Get:Latest")
	beego.Router("/top-rated", &web.OtherPageController{}, "Get:TopRatePage")
	beego.Router("/most-popular", &web.OtherPageController{}, "Get:MostViewedPage")

	// categories 页面
	beego.Router("/categories/?:category_id", &web.CategoryController{}, "Get:Index")

	// pornstars 页面
	beego.Router("/pornstars/?:pornstar_id", &web.PornstarController{}, "Get:Index")

	// pornstars 页面
	beego.Router("/tags/?:tag_id", &web.TagController{}, "Get:Index")

	// 搜索页面
	beego.Router("/search", &web.SearchController{})

	// 播放页面
	beego.Router("/play/:video_id", &web.PlayPageController{})

	beego.Router("/login", &web.LoginController{})
}
