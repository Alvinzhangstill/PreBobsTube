package web

import (
	"fmt"
	"github.com/MobileCPX/PreBobsTube/models/web"
)

type SearchController struct {
	BaseController
}

func (c *SearchController) Get() {
	var totalPage int
	search := c.GetString("s")
	currentPage := c.currentPageNum()
	c.setRandomTags()

	c.leftSidebarSort("most_views")
	c.Data["Videos"], totalPage = web.SearchVideos(c.Orm, search, "video_name", 24, currentPage)
	c.setPagination(currentPage, totalPage, "/search?s="+search+"&p=%d")
	c.Data["PageName"] = fmt.Sprintf(`Search results for " %s "`, search)
	c.setTpl("temp/pages/top_video.html")
}
