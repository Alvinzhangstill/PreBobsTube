package web

import (
	"bufio"
	"encoding/json"
	"github.com/MobileCPX/PreBobsTube/enums"
	"github.com/MobileCPX/PreBobsTube/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io"
	"os"
	"strconv"
)

//Videos 视频信息
type Videos struct {
	VideoID       string `json:"video_id" orm:"pk;column(id)"`
	Name          string `json:"video_name"`
	VideoPlayPage string `json:"play_page_url"`
	PreviewMP4    string `orm:"column(preview_mp4)"`
	VideoImageURL string `json:"image_url" orm:"column(video_image_url)"`
	VideoDH       string `json:"video_hd" orm:"column(video_dh)"`
	VideoDuration string `json:"video_duration"`
	LikeNum       int
	UnLikeNum     int
	LikeRate      int
	ViewNum       int
	UpLoadTime    string `json:"upload_time"`
	VideoURL      string
	Quality4K     string `json:"4k_url"`
	Quality1080P  string `json:"1080p_url"`
	Quality720P   string `json:"720p_url"`
	Quality480P   string `json:"480p_url"`
	Quality320P   string `json:"320p_url"`
	Quality240P   string `json:"240p_url"`

	CategoryIDs string `orm:"column(category_ids)" json:"category"`

	TagsIDs string `orm:"column(tag_ids)" json:"tags"`

	PornStarIDs  string `orm:"column(pornstar_ids)" json:"pornstar"`
	UpdateStauts bool   `orm:"column(update_status)"`
}

func InitVideoJOSNToTable() {
	o := orm.NewOrm()
	f, err := os.Open("source/video_data/4k_data.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		oneVideo := new(Videos)
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		err = json.Unmarshal([]byte(line), oneVideo)
		if err != nil {
			logs.Error("ReadVideoJOSN JSON 数据解析失败 ERROR: ", err.Error())
			//return
		}

		oneVideo.VideoDuration = oneVideo.VideoDuration + ":" + strconv.Itoa(util.GenerateRangeNum(1, 59))
		oneVideo.LikeNum, oneVideo.UnLikeNum, oneVideo.LikeRate = getLikeNumAndUnLikeNum(1000, 150)
		oneVideo.PornStarIDs = util.StrToSQLFilterList(oneVideo.PornStarIDs, "####")
		oneVideo.CategoryIDs = util.StrToSQLFilterList(oneVideo.CategoryIDs, "####")
		oneVideo.TagsIDs = util.StrToSQLFilterList(oneVideo.TagsIDs, "####")
		oneVideo.ViewNum = util.GenerateRangeNum(3000, 30000)
		oneVideo.GetPreviewMP4URL()
		_, err = o.Insert(oneVideo)
		if err != nil {
			logs.Error("ReadVideoJOSN 存入视频数据失败，error: ", err.Error())
		}
	}
}

// 返回两个参数  1：根据条件筛选出来的视频   2：该类视频的总页面
func VideosData(o orm.Ormer, sortType enums.VideoSortType, limitNum, pageNum int) (*[]Videos, int) {
	videos := new([]Videos)
	SQL := o.QueryTable(VideosTBName())
	switch sortType {
	case enums.VSMostView:
		SQL = SQL.OrderBy("-view_num")
	case enums.TopRateVideo:
		SQL = SQL.OrderBy("-like_rate")
	case enums.MostViewsVideo:
		SQL = SQL.OrderBy("-view_num")
	}

	countVideo, err := SQL.Count()
	totalPage := getTotalPage(int(countVideo), limitNum)

	if err != nil {
		logs.Error("VideosData 查询数据失败，ERROR: ", err.Error())
	}
	_, err = SQL.Limit(limitNum, (pageNum-1)*limitNum).All(videos)
	if err != nil {
		logs.Error("VideosData 查询数据失败，ERROR: ", err.Error())
	}
	videos = UpdateImageURL(*videos, "250")
	return videos, totalPage
}

func GetRandomVideos(o orm.Ormer, limitNum string) (*[]Videos, error) {
	videos := new([]Videos)
	SQL := "SELECT * FROM videos order by random() limit " + limitNum
	_, err := o.Raw(SQL).QueryRows(videos)
	if err != nil {
		logs.Error("GetRandomVideos 随机获取视频失败，ERROR: ", err.Error())
	}
	return videos, err
}
func (video *Videos) AddViewNum(o orm.Ormer) {
	video.ViewNum = video.ViewNum + util.GenerateRangeNum(5, 15)
}

// 获取播放页面预览图片
// minImageSize: 小尺寸图片大小   maxImageSize： 大尺寸图片大小
func (video *Videos) PreImageList(minImageSize, maxImageSize string, imageNum int) map[string]string {
	imageMap := make(map[string]string)
	videoIDIndex := []rune(video.VideoID)
	for i := 1; i <= imageNum; i++ {
		minImageURL := "http://cdnthumb5.spankbang.com/" + minImageSize + "/" + string(videoIDIndex[0]) + "/" +
			string(videoIDIndex[1]) + "/" + video.VideoID + "-t" + strconv.Itoa(i) + ".jpg"
		maxImageURL := "http://cdnthumb5.spankbang.com/" + maxImageSize + "/" + string(videoIDIndex[0]) + "/" +
			string(videoIDIndex[1]) + "/" + video.VideoID + "-t" + strconv.Itoa(i) + ".jpg"
		imageMap[minImageURL] = maxImageURL
	}
	return imageMap
}

// 获取视频好评率及点赞数
func (video *Videos) GetLikeRate() (likeRate, totalVotes int) {
	// 好评率  直接x100 用于模板渲染数据
	likeRate = video.LikeNum * 100 / (video.LikeNum + video.UnLikeNum)
	totalVotes = video.LikeNum + video.UnLikeNum
	return

}

// 获取可播放的视频URL
func (video *Videos) GetVideoMP4URL() *Videos {
	// 视频清晰度的优先级  720P 480P   320P 1080P  4K 320P
	if video.Quality720P != "" {
		video.VideoURL = video.Quality720P
	} else if video.Quality480P != "" {
		video.VideoURL = video.Quality480P
	} else if video.Quality320P != "" {
		video.VideoURL = video.Quality320P
	} else if video.Quality1080P != "" {
		video.VideoURL = video.Quality1080P
	} else if video.Quality4K != "" {
		video.VideoURL = video.Quality4K
	} else if video.Quality320P != "" {
		video.VideoURL = video.Quality320P
	}
	return video
}

// 通过videoID 查询视频信息
func (video *Videos) GetVideoByVideoID(o orm.Ormer, videoID string) error {
	video.VideoID = videoID
	err := o.Read(video)
	if err != nil {
		logs.Error("GetVideoByVideoID 通过videoID 查询视频信息失败，videoID: ", videoID)
	}
	return err
}

func getLikeNumAndUnLikeNum(maxTotalNum, minTotalNum int) (likeNum, unLikeNum, likeRate int) {
	totalNum := util.GenerateRangeNum(minTotalNum, maxTotalNum)
	likeRate = util.GenerateRangeNum(70, 100)
	likeNum = totalNum * likeRate / 100
	unLikeNum = totalNum - likeNum

	return
}

// getTotalPage 两个参数
// countVideo: 视频数量   onePageVideosNum： 单页视频数量
func getTotalPage(countVideo, onePageVideosNum int) int {
	totalPage := int(countVideo) / onePageVideosNum
	if int(countVideo)%onePageVideosNum != 0 {
		totalPage++
	}
	return totalPage
}

func (video *Videos) GetPreviewMP4URL() {
	videoIDIndex := []rune(video.VideoID)
	video.PreviewMP4 = "http://cdnthumb1.spankbang.com/0/" + string(videoIDIndex[0]) + "/" + string(videoIDIndex[1]) +
		"/" + video.VideoID + "-t.mp4"
}

func (video *Videos) GetVideoImageBySize(imageSize string) {
	videoIDs := []rune(video.VideoID)
	video.VideoImageURL = "http://cdnthumb5.spankbang.com/" + imageSize + "/" + string(videoIDs[0]) + "/" +
		string(videoIDs[1]) + "/" + video.VideoID + "-t6.jpg"

}
func UpdateImageURL(videos []Videos, imageSize string) *[]Videos {
	for i, oneVideo := range videos {
		videoIDs := []rune(oneVideo.VideoID)
		videos[i].VideoImageURL = "http://cdnthumb5.spankbang.com/" + imageSize + "/" + string(videoIDs[0]) + "/" + string(videoIDs[1]) + "/" + oneVideo.VideoID + "-t6.jpg"
	}
	return &videos
}


func SearchVideos(o orm.Ormer,searchKey,searchType string,limitNum ,pageNum int) (*[]Videos, int){
	videos := new([]Videos)
	SQL := o.QueryTable(VideosTBName())
	switch searchType {
	case "video_name":
		SQL = SQL.Filter("name__icontains",searchKey)
	case "category":
		SQL= SQL.Filter("category_ids__icontains",searchKey)
	case "tag":
		SQL = SQL.Filter("tag_ids__icontains",searchKey)
	case "pornstar":
		SQL = SQL.Filter("pornstar_ids__icontains",searchKey)
	}
	countVideo, err := SQL.Count()
	totalPage := getTotalPage(int(countVideo), limitNum)

	if err != nil {
		logs.Error("SearchVideos 查询数据失败，ERROR: ", err.Error())
	}
	_, err = SQL.Limit(limitNum, (pageNum-1)*limitNum).All(videos)
	if err != nil {
		logs.Error("SearchVideos 查询数据失败，ERROR: ", err.Error())
	}
	videos = UpdateImageURL(*videos, "250")
	return videos, totalPage
}