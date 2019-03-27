package web

import (
	"fmt"
	"github.com/MobileCPX/PreBobsTube/models/web"
)

type PornstarController struct {
	BaseController
}

// Most Viewed Videos 页面  最多观看页面
func (c *PornstarController) Index() {
	var totalPage int
	pornstarID := c.Ctx.Input.Param(":pornstar_id")
	c.setRandomTags()
	c.leftSidebarSort("pornstar")
	c.Data["nav_selected"] = "mmodels"
	if pornstarID == "" {
		c.Data["Pornstars"], _ = web.GetAllVideoPornstars(c.Orm)
		c.setTpl("temp/pages/pornstar.html")
	} else {
		pornstar := new(web.PornStars)
		_ = pornstar.GetPornstarByPornstarID(c.Orm, pornstarID)

		currentPage := c.currentPageNum()
		//c.Data["Videos"], totalPage = web.VideosData(c.Orm, enums.MostViewsVideo, 24, currentPage)
		c.Data["Pornstar"] = pornstar
		c.Data["Videos"], totalPage = web.SearchVideos(c.Orm, pornstarID, "pornstar", 24, currentPage)
		c.setPagination(currentPage, totalPage, "/pornstars/"+pornstarID+"?p=%d")
		c.Data["PageName"] = fmt.Sprintf(`Search results for " %s "`, pornstar.PornStarName)
		c.setTpl("temp/pages/pornstar_info.html")
	}
}
