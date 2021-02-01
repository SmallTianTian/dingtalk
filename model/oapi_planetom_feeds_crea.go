package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPlanetomFeedsCreateRequest() *OapiPlanetomFeedsCreateRequest {
	return &OapiPlanetomFeedsCreateRequest{}
}

type OapiPlanetomFeedsCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiPlanetomFeedsCreateResponse
	AppointBeginTime   int64
	CoverUrl           string
	FeedType           int64
	GroupIds           string
	Introduction       string
	PicIntroductionUrl string
	Title              string
	TopHttpMethod      string
	TopResponseType    string
	Userid             string
}

func (this *OapiPlanetomFeedsCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPlanetomFeedsCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPlanetomFeedsCreateRequest) SetAppointBeginTime(appointBeginTime2 int64) {
	this.AppointBeginTime = appointBeginTime2
}
func (this *OapiPlanetomFeedsCreateRequest) GetAppointBeginTime() int64 {
	return this.AppointBeginTime
}
func (this *OapiPlanetomFeedsCreateRequest) SetCoverUrl(coverUrl2 string) {
	this.CoverUrl = coverUrl2
}
func (this *OapiPlanetomFeedsCreateRequest) GetCoverUrl() string {
	return this.CoverUrl
}
func (this *OapiPlanetomFeedsCreateRequest) SetFeedType(feedType2 int64) {
	this.FeedType = feedType2
}
func (this *OapiPlanetomFeedsCreateRequest) GetFeedType() int64 {
	return this.FeedType
}
func (this *OapiPlanetomFeedsCreateRequest) SetGroupIds(groupIds2 string) {
	this.GroupIds = groupIds2
}
func (this *OapiPlanetomFeedsCreateRequest) GetGroupIds() string {
	return this.GroupIds
}
func (this *OapiPlanetomFeedsCreateRequest) SetIntroduction(introduction2 string) {
	this.Introduction = introduction2
}
func (this *OapiPlanetomFeedsCreateRequest) GetIntroduction() string {
	return this.Introduction
}
func (this *OapiPlanetomFeedsCreateRequest) SetPicIntroductionUrl(picIntroductionUrl2 string) {
	this.PicIntroductionUrl = picIntroductionUrl2
}
func (this *OapiPlanetomFeedsCreateRequest) GetPicIntroductionUrl() string {
	return this.PicIntroductionUrl
}
func (this *OapiPlanetomFeedsCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiPlanetomFeedsCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiPlanetomFeedsCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiPlanetomFeedsCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiPlanetomFeedsCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.planetom.feeds.create"
}
func (this *OapiPlanetomFeedsCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPlanetomFeedsCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPlanetomFeedsCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPlanetomFeedsCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPlanetomFeedsCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["appoint_begin_time"] = this.AppointBeginTime
	txtParams["cover_url"] = this.CoverUrl
	txtParams["feed_type"] = this.FeedType
	txtParams["group_ids"] = this.GroupIds
	txtParams["introduction"] = this.Introduction
	txtParams["pic_introduction_url"] = this.PicIntroductionUrl
	txtParams["title"] = this.Title
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiPlanetomFeedsCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppointBeginTime, "appointBeginTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPlanetomFeedsCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPlanetomFeedsCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPlanetomFeedsCreateResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
