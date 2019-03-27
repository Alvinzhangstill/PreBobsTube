package web

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/MobileCPX/PreBobsTube/enums"
	"github.com/MobileCPX/PreBobsTube/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
)

type Category struct {
	CategoryID       string `orm:"pk;column(category_id)"`
	CategoryName     string `json:"CategoryName"`
	CategoryImageURL string `json:"CategoryImageURL"`
	Score            float32
	LikeRate         int
	VideoNum         int
}

func InitCategoryJOSNToTable() {
	o := orm.NewOrm()
	f, err := os.Open("source/video_data/category_data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// var datas []string
	rd := bufio.NewReader(f)
	for {
		var oneCategory Category
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		json.Unmarshal([]byte(line), &oneCategory)
		oneCategory.VideoNum = util.GenerateRangeNum(500, 1000)
		oneCategory.Score = util.GenerateFloatNum(5, 2, 1)
		oneCategory.LikeRate = util.GenerateRangeNum(50, 100)
		oneCategory.CategoryID = util.StrToStrid(oneCategory.CategoryName)
		o.Insert(&oneCategory)
	}
}

func RandomCategories(o orm.Ormer, categoryNum int) (*[]Category, error) {
	tags := new([]Category)
	SQL := fmt.Sprintf("SELECT * FROM category ORDER BY random() LIMIT %d", categoryNum)
	_, err := o.Raw(SQL).QueryRows(tags)
	if err != nil {
		logs.Error("RandomCategors 获取tags失败, ERROR: ", err.Error())
	}
	return tags, err
}

// 根据评分排序获取Categories信息
func TopCategories(o orm.Ormer, categoryNum int) (*[]Category, error) {
	categories := new([]Category)
	_, err := o.QueryTable(CategoryTBName()).OrderBy("-score").Limit(categoryNum).All(categories)
	if err != nil {
		logs.Error("TopCategories 获取categories失败, ERROR: ", err.Error())
	}
	return categories, err
}

func GetCategoryByTagidList(o orm.Ormer, categoryIDList string) (*[]Category, error) {
	categories := new([]Category)
	SQL := "SELECT * FROM category where category_id in  " + categoryIDList
	_, err := o.Raw(SQL).QueryRows(categories)
	if err != nil {
		logs.Error("GetTagsByTagidList 通过categoryId list 查询Category 信息失败，ERROR:", err.Error())
	}
	return categories, err
}

func GetTopCategory(o orm.Ormer, limitNum string) (*[]LeftSidebarData, error) {
	sidebarData := new([]LeftSidebarData)
	SQL := "SELECT category_id as sidebar_i_d,category_name as sidebar_name,score as sidebar_score from" +
		" category order by score desc limit " + limitNum
	_, err := o.Raw(SQL).QueryRows(sidebarData)
	if err != nil {
		logs.Error("GetTopCategory 查询Category 信息失败，ERROR:", err.Error())
	}
	return sidebarData, err
}

func GetAllVideoCategoies(o orm.Ormer, sortType enums.CategorySortType) (*[]Category, error) {
	categories := new([]Category)
	SQL := o.QueryTable(CategoryTBName())
	switch sortType {
	case enums.Alphabetically:
		SQL = SQL.OrderBy("category")
	case enums.TopRateCategory:
		SQL = SQL.OrderBy("-like_rate")
	case enums.MostViewCategory:
		SQL = SQL.OrderBy("-score")
	case enums.MostVidesCategory:
		SQL = SQL.OrderBy("-video_num")
	}
	_, err := SQL.All(categories)
	if err != nil {
		logs.Error("GetAllVideoCategoies 获取所有Categoies失败，ERROR:", err.Error())
	}
	return categories, err
}

func (category *Category) GetCategoryByCategoryID(o orm.Ormer, categoryID string) error {
	err := o.QueryTable(CategoryTBName()).Filter("category_id", categoryID).One(category)
	if err != nil {
		logs.Error("GetCategoryByCategoryID 查询Category 信息失败，ERROR:", err.Error())
	}
	return err
}
