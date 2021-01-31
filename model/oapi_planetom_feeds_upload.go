package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPlanetomFeedsUploadRequest() *OapiPlanetomFeedsUploadRequest {
	return &OapiPlanetomFeedsUploadRequest{}
}

type OapiPlanetomFeedsUploadRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPlanetomFeedsUploadResponse
	FeedAppId       int64
	FeedInfoModels  string
	TopHttpMethod   string
	TopResponseType string
	UserPhone       string
}

func (this *OapiPlanetomFeedsUploadRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPlanetomFeedsUploadRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPlanetomFeedsUploadRequest) SetFeedAppId(feedAppId2 int64) {
	this.FeedAppId = feedAppId2
}
func (this *OapiPlanetomFeedsUploadRequest) GetFeedAppId() int64 {
	return this.FeedAppId
}
func (this *OapiPlanetomFeedsUploadRequest) SetFeedInfoModels(feedInfoModels2 string) {
	this.FeedInfoModels = feedInfoModels2
}
func (this *OapiPlanetomFeedsUploadRequest) GetFeedInfoModels() string {
	return this.FeedInfoModels
}
func (this *OapiPlanetomFeedsUploadRequest) SetUserPhone(userPhone2 string) {
	this.UserPhone = userPhone2
}
func (this *OapiPlanetomFeedsUploadRequest) GetUserPhone() string {
	return this.UserPhone
}
func (this *OapiPlanetomFeedsUploadRequest) GetApiMethodName() string {
	return "dingtalk.oapi.planetom.feeds.upload"
}
func (this *OapiPlanetomFeedsUploadRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPlanetomFeedsUploadRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPlanetomFeedsUploadRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPlanetomFeedsUploadRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPlanetomFeedsUploadRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["feed_app_id"] = this.FeedAppId
	txtParams["feed_info_models"] = this.FeedInfoModels
	txtParams["user_phone"] = this.UserPhone
	return txtParams
}
func (this *OapiPlanetomFeedsUploadRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FeedAppId, "feedAppId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPlanetomFeedsUploadRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPlanetomFeedsUploadRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FeedContentModel struct {
	FeedType int64  `json:"feed_type,omitempty"`
	Title    string `json:"title,omitempty"`
	VideoUrl string `json:"video_url,omitempty"`
}
type FeedBaseModel struct {
	CoverUrl        string `json:"cover_url,omitempty"`
	CustTag         string `json:"cust_tag,omitempty"`
	Introduction    string `json:"introduction,omitempty"`
	PicIntroduction string `json:"pic_introduction,omitempty"`
	Title           string `json:"title,omitempty"`
}
type FeedPayModel struct {
	CsPhone      string `json:"cs_phone,omitempty"`
	DisEndTime   int64  `json:"dis_end_time,omitempty"`
	DisPrice     int64  `json:"dis_price,omitempty"`
	DisStartTime int64  `json:"dis_start_time,omitempty"`
	LimitCount   int64  `json:"limit_count,omitempty"`
	NeedPay      bool   `json:"need_pay,omitempty"`
	Preferential bool   `json:"preferential,omitempty"`
	Price        int64  `json:"price,omitempty"`
}
type BatchUploadFeedInfoModel struct {
	BaseModel     FeedBaseModel      `json:"base_model,omitempty"`
	ContentModels []FeedContentModel `json:"content_models,omitempty"`
	PayModel      FeedPayModel       `json:"pay_model,omitempty"`
	ResourceType  int64              `json:"resource_type,omitempty"`
	WorkGroupName string             `json:"work_group_name,omitempty"`
}
type OapiPlanetomFeedsUploadResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
