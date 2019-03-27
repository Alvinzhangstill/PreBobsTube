package web

import (
	"github.com/MobileCPX/PreBobsTube/enums"
	"github.com/MobileCPX/PreBobsTube/models/web"
)

type OtherPageController struct {
	BaseController
}

// 最新视频页面
func (c *OtherPageController) Latest() {
	var totalPage int
	currentPage := c.currentPageNum()
	c.Data["PageName"] = "New Videos"
	c.Data["nav_selected"] = "mmovies"
	c.setRandomTags()
	c.leftSidebarSort("last")
	c.Data["Videos"], totalPage = web.VideosData(c.Orm, enums.VSNewVideo, 24, currentPage)

	c.setPagination(currentPage, totalPage, "/latest-updates?p=%d")
	c.setTpl("temp/pages/top_video.html")
}

// TOP RATE 页面
func (c *OtherPageController) TopRatePage() {
	var totalPage int
	currentPage := c.currentPageNum()
	c.Data["nav_selected"] = "mtoprated"
	c.setRandomTags()
	c.leftSidebarSort("top_rate")
	c.Data["Videos"], totalPage = web.VideosData(c.Orm, enums.TopRateVideo, 24, currentPage)
	c.Data["PageName"] = "Top Rated Videos"
	c.setPagination(currentPage, totalPage, "/top-rated?p=%d")
	c.setTpl("temp/pages/top_video.html")
}

// Most Viewed Videos 页面  最多观看页面
func (c *OtherPageController) MostViewedPage() {
	var totalPage int
	currentPage := c.currentPageNum()
	c.Data["nav_selected"] = "mmostpopular"
	c.setRandomTags()
	c.leftSidebarSort("most_views")
	c.Data["Videos"], totalPage = web.VideosData(c.Orm, enums.MostViewsVideo, 24, currentPage)
	c.Data["PageName"] = "Most Viewed Videos"
	c.setPagination(currentPage, totalPage, "/most-popular?p=%d")
	c.setTpl("temp/pages/top_video.html")
}
