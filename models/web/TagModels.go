package web

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/MobileCPX/PreBobsTube/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
)

type Tag struct {
	TagID   string `orm:"pk;column(tag_id)"`
	TagName string `orm:"column(tag_name" json:"tags"`
	Score   float32
	View    int
}

// InitTagsJSONToTable 初始化Tags 数据,配置文件中读取数据
func InitTagsJSONToTable() {
	o := orm.NewOrm()
	f, err := os.Open("source/video_data/tag.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// var datas []string
	rd := bufio.NewReader(f)
	for {
		var oneTag Tag
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		err = json.Unmarshal([]byte(line), &oneTag)
		if err != nil {
			logs.Error("InitTagsJSONToTable json 数据解析失败，ERROR: ", err.Error())
		}
		oneTag.View = util.GenerateRangeNum(400000, 800000)
		oneTag.TagID = util.StrToStrid(oneTag.TagName)
		oneTag.Score = util.GenerateFloatNum(5, 2, 1)
		_, err = o.Insert(&oneTag)
		if err != nil {
			logs.Error("新插入tag失败，ERROR: ", err.Error())
		}
	}
}

func RandomTags(o orm.Ormer, tagNum int) (*[]Tag, error) {
	tags := new([]Tag)
	SQL := fmt.Sprintf("SELECT * FROM tag ORDER BY random() LIMIT %d", tagNum)
	_, err := o.Raw(SQL).QueryRows(tags)
	if err != nil {
		logs.Error("RandomTags 获取tags失败, ERROR: ", err.Error())
	}
	return tags, err
}

func GetTagsByTagidList(o orm.Ormer, tagIDList string) (*[]Tag, error) {
	tags := new([]Tag)
	SQL := "SELECT * FROM tag where tag_id in  " + tagIDList
	_, err := o.Raw(SQL).QueryRows(tags)
	if err != nil {
		logs.Error("GetTagsByTagidList 通过tagsId list 查询Tags 信息失败，ERROR:", err.Error())
	}
	return tags, err
}

func GetTopTags(o orm.Ormer, limitNum string) (*[]LeftSidebarData, error) {
	sidebarData := new([]LeftSidebarData)
	SQL := "SELECT tag_id as sidebar_i_d,tag_name as sidebar_name,score as sidebar_score from" +
		" tag order by score desc limit " + limitNum
	_, err := o.Raw(SQL).QueryRows(sidebarData)
	if err != nil {
		logs.Error("GetTopTags 查询Category 信息失败，ERROR:", err.Error())
	}
	return sidebarData, err
}

func (tag *Tag) GetTagByTagID(o orm.Ormer, tagID string) error {
	err := o.QueryTable(TagTBName()).Filter("tag_id", tagID).One(tag)
	if err != nil {
		logs.Error("GetTagByTagID 查询Tag 信息失败，ERROR:", err.Error())
	}
	return err
}
