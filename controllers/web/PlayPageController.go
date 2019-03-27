package web

import (
	"github.com/MobileCPX/PreBobsTube/models/web"
)

type PlayPageController struct {
	BaseController
}

func (c *PlayPageController) Get() {
	videoID := c.Ctx.Input.Param(":video_id")
	playVideo := new(web.Videos)
	err := playVideo.GetVideoByVideoID(c.Orm, videoID)
	if err != nil {
		c.redirect("/404")
	}

	playVideo.GetVideoMP4URL() // 获取可播放视频链接
	playVideo.GetVideoImageBySize("720")

	// 视频的相关Tags，Categories，Pornstars
	c.Data["VideoTags"], _ = web.GetTagsByTagidList(c.Orm, playVideo.TagsIDs)
	c.Data["VideoCategories"], _ = web.GetCategoryByTagidList(c.Orm, playVideo.CategoryIDs)
	c.Data["VideoPornstars"], _ = web.GetPornstarsByTagidList(c.Orm, playVideo.PornStarIDs)

	// 相关视频
	c.Data["RelatedVideos"], _ = web.GetRandomVideos(c.Orm, "15")
	c.Data["PlayVideo"] = playVideo
	c.Data["PreImageList"] = playVideo.PreImageList("250", "2000", 10)
	c.Data["LikeRate"], c.Data["TotalVotes"] = playVideo.GetLikeRate()
	playVideo.AddViewNum(c.Orm)
	c.setTpl("temp/pages/play_page.html")
}
