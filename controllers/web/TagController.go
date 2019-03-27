package web

import (
	"fmt"
	"github.com/MobileCPX/PreBobsTube/models/web"
)

type TagController struct {
	BaseController
}

// Most Viewed Videos 页面  最多观看页面
func (c *TagController) Index() {
	tagID := c.Ctx.Input.Param(":tag_id")
	c.setRandomTags()
	var totalPage int
	c.leftSidebarSort("category")

	tag := new(web.Tag)
	_ = tag.GetTagByTagID(c.Orm, tagID)

	currentPage := c.currentPageNum()
	c.leftSidebarSort("most_views")
	c.Data["Videos"], totalPage = web.SearchVideos(c.Orm, tagID, "tag", 24, currentPage)
	c.setPagination(currentPage, totalPage, "/tags/"+tagID+"?p=%d")
	c.Data["PageName"] = fmt.Sprintf(`Search results for " %s "`, tag.TagName)
	c.setTpl("temp/pages/top_video.html")

}
