package web

import (
	"bufio"
	"encoding/json"
	"github.com/MobileCPX/PreBobsTube/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
)

type PornStars struct {
	PornStarID    string `orm:"pk;column(pornstar_id)"`
	PornStarName  string `json:"pornstar_name"`
	PornStarImage string `json:"PornstarImage"`
	Score         float32
	LikeRate      int
	ViewNum       int
	VideoNum      int
}

func InitPornStarsJOSNToTable() {
	o := orm.NewOrm()
	f, err := os.Open("source/video_data/pornstars_data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// var datas []string
	rd := bufio.NewReader(f)
	for {
		var pornStar PornStars
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		json.Unmarshal([]byte(line), &pornStar)
		// fmt.Println(line)
		pornStar.ViewNum = util.GenerateRangeNum(1000, 4000)
		pornStar.PornStarID = util.StrToStrid(pornStar.PornStarName)
		pornStar.Score = util.GenerateFloatNum(5, 2, 1)
		pornStar.LikeRate = util.GenerateRangeNum(50, 100)
		o.Insert(&pornStar)
	}

}

func GetPornstarsByTagidList(o orm.Ormer, pornstarIDList string) (*[]PornStars, error) {
	pornstars := new([]PornStars)
	SQL := "SELECT * FROM porn_stars where pornstar_id in  " + pornstarIDList
	_, err := o.Raw(SQL).QueryRows(pornstars)
	if err != nil {
		logs.Error("GetTagsByTagidList 通过pornstarId list 查询Pornstars 信息失败，ERROR:", err.Error())
	}
	return pornstars, err
}

func GetTopPornstars(o orm.Ormer, limitNum string) (*[]LeftSidebarData, error) {
	sidebarData := new([]LeftSidebarData)
	SQL := "SELECT pornstar_id as sidebar_i_d,porn_star_name as sidebar_name,score as sidebar_score from" +
		" porn_stars order by score desc limit " + limitNum
	_, err := o.Raw(SQL).QueryRows(sidebarData)
	if err != nil {
		logs.Error("GetTopPornstars 查询Category 信息失败，ERROR:", err.Error())
	}
	return sidebarData, err
}

func (pornstar *PornStars) GetPornstarByPornstarID(o orm.Ormer, pornstarID string) error {
	err := o.QueryTable(PornstarTBName()).Filter("pornstar_id", pornstarID).One(pornstar)
	if err != nil {
		logs.Error("GetPornstarByPornstarID 查询Pornstar 信息失败，ERROR:", err.Error())
	}
	return err
}

func GetAllVideoPornstars(o orm.Ormer) (*[]PornStars, error) {
	pornstars := new([]PornStars)
	SQL := o.QueryTable(PornstarTBName())

	//switch sortType {
	//case enums.Alphabetically:
	//	SQL = SQL.OrderBy("category")
	//case enums.TopRateCategory:
	//	SQL = SQL.OrderBy("-like_rate")
	//case enums.MostViewCategory:
	//	SQL = SQL.OrderBy("-score")
	//case enums.MostVidesCategory:
	//	SQL = SQL.OrderBy("-video_num")
	//}

	_, err := SQL.OrderBy("-score").All(pornstars)
	if err != nil {
		logs.Error("GetAllVideoCategoies 获取所有Categoies失败，ERROR:", err.Error())
	}
	return pornstars, err
}
