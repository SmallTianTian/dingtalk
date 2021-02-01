package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCirclePostListRequest() *OapiEduCirclePostListRequest {
	return &OapiEduCirclePostListRequest{}
}

type OapiEduCirclePostListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiEduCirclePostListResponse
	OpenFeedQueryParam string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiEduCirclePostListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCirclePostListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCirclePostListRequest) SetOpenFeedQueryParam(openFeedQueryParam2 string) {
	this.OpenFeedQueryParam = openFeedQueryParam2
}
func (this *OapiEduCirclePostListRequest) GetOpenFeedQueryParam() string {
	return this.OpenFeedQueryParam
}
func (this *OapiEduCirclePostListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.circle.post.list"
}
func (this *OapiEduCirclePostListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCirclePostListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCirclePostListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCirclePostListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCirclePostListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_feed_query_param"] = this.OpenFeedQueryParam
	return txtParams
}
func (this *OapiEduCirclePostListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCirclePostListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCirclePostListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenFeedQueryParam struct {
	BizType   int64  `json:"biz_type,omitempty"`
	ClassId   int64  `json:"class_id,omitempty"`
	Count     int64  `json:"count,omitempty"`
	Cursor    int64  `json:"cursor,omitempty"`
	FeedType  int64  `json:"feed_type,omitempty"`
	StudentId string `json:"student_id,omitempty"`
	TopicId   int64  `json:"topic_id,omitempty"`
	UserRole  string `json:"user_role,omitempty"`
	Userid    string `json:"userid,omitempty"`
}
type OapiEduCirclePostListResponse struct {
	taobao.TaobaoResponse
	Result  OpenCircleTopicResponse `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type OrgUserDto struct {
	ShowName string `json:"show_name,omitempty"`
	StaffId  string `json:"staff_id,omitempty"`
}
type Comments struct {
	CommentId  int64      `json:"comment_id,omitempty"`
	Content    string     `json:"content,omitempty"`
	OriginUser OrgUserDto `json:"origin_user,omitempty"`
}
type Author struct {
	AvatarMediaId string `json:"avatar_media_id,omitempty"`
	IconMediaId   string `json:"icon_media_id,omitempty"`
	IsOwner       bool   `json:"is_owner,omitempty"`
	Nick          string `json:"nick,omitempty"`
	Owner         bool   `json:"owner,omitempty"`
	ShowName      string `json:"show_name,omitempty"`
	StaffId       string `json:"staff_id,omitempty"`
	Tag           int64  `json:"tag,omitempty"`
	Title         string `json:"title,omitempty"`
	Type          string `json:"type,omitempty"`
	UserRole      string `json:"user_role,omitempty"`
}
type Content struct {
	ContentType int64  `json:"content_type,omitempty"`
	GeoContent  string `json:"geo_content,omitempty"`
	Text        string `json:"text,omitempty"`
}
type Posts struct {
	Author   Author     `json:"author,omitempty"`
	BizId    string     `json:"biz_id,omitempty"`
	Comments []Comments `json:"comments,omitempty"`
	Content  Content    `json:"content,omitempty"`
	CreateAt int64      `json:"create_at,omitempty"`
	FeedType int64      `json:"feed_type,omitempty"`
	PostId   int64      `json:"post_id,omitempty"`
	Status   int64      `json:"status,omitempty"`
	Tags     string     `json:"tags,omitempty"`
}
type OpenCircleTopicResponse struct {
	HasMore bool    `json:"has_more,omitempty"`
	Posts   []Posts `json:"posts,omitempty"`
}
