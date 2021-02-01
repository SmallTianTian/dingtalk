package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPlanetomFeedsStatisticGetRequest() *OapiPlanetomFeedsStatisticGetRequest {
	return &OapiPlanetomFeedsStatisticGetRequest{}
}

type OapiPlanetomFeedsStatisticGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPlanetomFeedsStatisticGetResponse
	FeedId          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiPlanetomFeedsStatisticGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPlanetomFeedsStatisticGetRequest) SetFeedId(feedId2 string) {
	this.FeedId = feedId2
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetFeedId() string {
	return this.FeedId
}
func (this *OapiPlanetomFeedsStatisticGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.planetom.feeds.statistic.get"
}
func (this *OapiPlanetomFeedsStatisticGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPlanetomFeedsStatisticGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPlanetomFeedsStatisticGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["feed_id"] = this.FeedId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiPlanetomFeedsStatisticGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FeedId, "feedId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPlanetomFeedsStatisticGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPlanetomFeedsStatisticGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenFeedInfoModel `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type OpenFeedInfoModel struct {
	ChatIds      []string  `json:"chat_ids,omitempty"`
	CoverUrl     string    `json:"cover_url,omitempty"`
	FeedId       string    `json:"feed_id,omitempty"`
	FeedType     int64     `json:"feed_type,omitempty"`
	Introduction string    `json:"introduction,omitempty"`
	JumpUrl      string    `json:"jump_url,omitempty"`
	StartTime    time.Time `json:"start_time,omitempty"`
	Status       int64     `json:"status,omitempty"`
	SubStatus    int64     `json:"sub_status,omitempty"`
	Title        string    `json:"title,omitempty"`
	Userid       int64     `json:"userid,omitempty"`
}
