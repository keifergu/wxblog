package model

import (
	"time"
)

type (
	article struct {
		Id         int               `json:"id"`
		Title      string            `json:"title" form:"title"`
		Content    string            `json:"content" form:"title"`
		CreateTime int64             `json:"create_time"`
		UpdateTime int64             `json:"update_time"`
		ClickCount int               `json:"click_count"`
		LikeReader map[string]string `json:"reader_like"`
		Comments   []*comment        `json:"comments"`
		TopicId    int               `json:"topic_id"`
	}
	comment struct {
		Id         int    `json:"id"`
		Content    string `json:"content" form:"content"`
		CreateTime int64  `json:"create_time"`
		NickName   string `json:"nick_name" form:"nick_name"`
		AvatarUrl  string `json:"avatar_url"`
		ArticleId  int    `json:"article_id"`
		//Reader *reader `json:"reader"`
	}
	topic struct {
		Id         int        `json:"id"`
		TopicName  string     `json:"topic_name" form:"topic_name"`
		LikeCount  int        `json:"like_count"`
		CreateTime int64      `json:"create_time"`
		Articles   []*article `json:"articles"`
	}
	//reader struct {
	//	Id int `json:"id"`
	//	NickName string `json:"nick_name"`
	//	AvatarUrl string `json:"avatar_url"`
	//	Gender string `json:"gender"` //性别 0：未知、1：男、2：女
	//	Province string `json:"province"`
	//	City string `json:"city"`
	//	Country string `json:"country"`
	//	CommentId int `json:"comment_id"`
	//}
)

func NewArticle(id, topicid int, title, content string) *article {
	t := time.Now().Unix()
	return &article{
		Id:         id,
		Title:      title,
		Content:    content,
		CreateTime: t,
		UpdateTime: t,
		LikeReader: make(map[string]string),
		Comments:   []*comment{},
		TopicId:    topicid,
	}
}

func NewArticleSlice() []*article {
	return []*article{}
}

func NewComment(id, articleid int, content, nickname, avatarurl string) *comment {
	t := time.Now().Unix()
	return &comment{
		Id:         id,
		Content:    content,
		CreateTime: t,
		NickName:   nickname,
		AvatarUrl:  avatarurl,
		ArticleId:  articleid,
	}
}

func NewTopic(id int, topicname string) *topic {
	t := time.Now().Unix()
	return &topic{
		Id:         id,
		TopicName:  topicname,
		CreateTime: t,
		Articles:   []*article{},
	}
}

//func NewReader(nickname, avatarurl, province, city, country string, gender int) *reader {
//	g := ""
//	switch gender {
//	case 1:
//		g = "男"
//	case 2:
//		g = "女"
//	default:
//		g = "未知"
//	}
//	return &reader{
//		NickName: nickname,
//		AvatarUrl: avatarurl,
//		Gender: g,
//		Province: province,
//		City: city,
//		Country: country,
//	}
//}
