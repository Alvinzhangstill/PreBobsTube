package web

import (
	"fmt"
	"github.com/MobileCPX/PreBobsTube/enums"
	"github.com/MobileCPX/PreBobsTube/models/web"
)

type CategoryController struct {
	BaseController
}

// Most Viewed Videos 页面  最多观看页面
func (c *CategoryController) Index() {
	categoryID := c.Ctx.Input.Param(":category_id")
	c.setRandomTags()
	var totalPage int
	c.leftSidebarSort("category")
	c.Data["nav_selected"] = "mcategories"
	if categoryID == "" {
		if c.Ctx.Input.IsAjax() {

		}

		c.Data["Categories"], _ = web.GetAllVideoCategoies(c.Orm, enums.TopRateCategory)
		c.setTpl("temp/pages/category.html")
	} else {

		category := new(web.Category)
		_ = category.GetCategoryByCategoryID(c.Orm, categoryID)

		currentPage := c.currentPageNum()
		c.leftSidebarSort("most_views")
		c.Data["Videos"], totalPage = web.SearchVideos(c.Orm, categoryID, "category", 24, currentPage)
		c.setPagination(currentPage, totalPage, "/categories/"+categoryID+"?p=%d")
		c.Data["PageName"] = fmt.Sprintf(`Search results for " %s "`, category.CategoryName)
		c.setTpl("temp/pages/top_video.html")
	}

}

// category 功能先不要
//func (c *CategoryController) CategorySortReq() {
//	sortType := enums.DefaultSort
//	if c.Ctx.Input.IsAjax() {
//		sortStr := c.GetString("sort")
//		switch sortStr {
//		case "Alphabetically":
//			sortType = enums.Alphabetically
//		case "TopRated":
//			sortType = enums.TopRateCategory
//		case "MostViewed":
//			sortType = enums.MostViewCategory
//		case "MostVideos":
//			sortType = enums.MostVidesCategory
//		}
//		c.Data["Categories"], _ = web.GetAllVideoCategoies(c.Orm, sortType)
//	}
//	c.Data["Categories"], _ = web.GetAllVideoCategoies(c.Orm, enums.TopRateCategory)
//}
//
//func (c *CategoryController) returnHTML(categories *[]web.Category) {
//	titleName := "Categories Having Most Videos"
//	categoriesItem := ""
//	for _, oneCategory := range *categories {
//		categoriesItem += fmt.Sprintf(`
//                    <a class="item" href="/categories/%s/" title="%s">
//                        <div class="img">
//                            <img class="thumb" src="%s"
//                                 alt="%s"/>
//                        </div>
//                        <strong class="title">%s</strong>
//                        <div class="wrap">
//                            <div class="videos">%d videos</div>
//
//                            <div class="rating positive">
//                                %d%
//                            </div>
//                        </div>
//                    </a>`, oneCategory.CategoryID, oneCategory.CategoryName, oneCategory.CategoryImageURL,
//			oneCategory.CategoryName, oneCategory.CategoryName, oneCategory.VideoNum, oneCategory.LikeRate)
//	}
//	html := fmt.Sprintf(`<div id="list_categories_categories_list">
//        <div class="headline">
//            <h2>
//                %s </h2>
//
//            <div class="sort">
//                <span class="icon type-sort"></span>
//                <strong>Most Videos</strong>
//                <ul id="list_categories_categories_list_sort_list">
//                    <li>
//                        <a data-action="ajax" data-container-id="list_categories_categories_list_sort_list"
//                           data-block-id="list_categories_categories_list" data-parameters="sort_by:title">Alphabetically</a>
//                    </li>
//                    <li>
//                        <a data-action="ajax" data-container-id="list_categories_categories_list_sort_list"
//                           data-block-id="list_categories_categories_list"
//                           data-parameters="sort_by:avg_videos_popularity">Most Viewed</a>
//                    </li>
//                    <li>
//                        <a data-action="ajax" data-container-id="list_categories_categories_list_sort_list"
//                           data-block-id="list_categories_categories_list" data-parameters="sort_by:avg_videos_rating">Top
//                            Rated</a>
//                    </li>
//                </ul>
//            </div>
//
//            <div class="sort">
//                <span class="icon type-video"></span>
//                <strong>Videos</strong>
//                <ul>
//                    <li><a href="https://bobs-tube.com/albums/categories/">Albums</a></li>
//                </ul>
//            </div>
//        </div>
//        <div class="box">
//            <div class="list-categories">
//                <div class="margin-fix" id="list_categories_categories_list_items">
//
//
//
//                </div>
//            </div>
//        </div>
//    </div>`,titleName,)
//
//
//}
