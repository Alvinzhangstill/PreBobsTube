package web

import (
	"github.com/MobileCPX/PreBobsTube/enums"
	m "github.com/MobileCPX/PreBobsTube/models/web"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	var totalPage int
	currentPage := c.currentPageNum()
	c.Data["nav_selected"] = "mhome"

	//返回两个参数  1：根据条件筛选出来的视频   2：该类视频的总页面
	c.Data["mostViewsVideos"], _ = m.VideosData(c.Orm, enums.VSMostView, 10, 1)
	c.Data["newVideos"], totalPage = m.VideosData(c.Orm, enums.VSNewVideo, 20, currentPage)
	c.setRandomTags()
	c.setPagination(currentPage, totalPage, "/?p=%d")

	// 如果是第一页就显示mostViewsVideos ，其他页面隐藏
	c.Data["currentPage"] = currentPage

	c.setTpl("temp/pages/home.html")
}
