package web

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Language 页面语言适配
type Language struct {
	Language         string `orm:"pk;column(language)" json:"language"`
	Home             string `json:"home"`
	Newest           string `json:"newest"`
	NewestVideos     string `json:"newest_videos"`
	Popular          string `json:"popular"`
	History          string `json:"history"`
	RecentlyUpdated  string `json:"recently_updated"`
	PopularAnime     string `json:"popular_anime"`
	RelatedVideo     string `json:"related_video"`
	LoadMoreVideos   string `json:"load_more_videos"`
	Refresh          string `json:"refresh"`
	Subscription     string `json:"subscription"`
	Likes            string `json:"likes"`
	SelectEpisodes   string `json:"select_episodes"`
	Views            string `json:"views"`
	RelatedRecommend string `json:"related_recommend"`
	Login            string
}

func (language *Language)LanguageAdaption(lang string, o orm.Ormer) (*Language,error){
	err := o.QueryTable(LanguageTBName()).Filter("language", lang).One(language)
	if err != nil {
		logs.Error("LanguageAdaption 未查询到", lang, "国家语言", err.Error())
		err := o.QueryTable(LanguageTBName()).Filter("language", "en").One(language)
		if err != nil {
			logs.Error("LanguageAdaption 没有英语语言", err.Error())
		}
	}
	return language,err
}
