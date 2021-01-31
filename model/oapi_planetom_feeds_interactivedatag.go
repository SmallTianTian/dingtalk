package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPlanetomFeedsInteractivedataGetRequest() *OapiPlanetomFeedsInteractivedataGetRequest {
	return &OapiPlanetomFeedsInteractivedataGetRequest{}
}

type OapiPlanetomFeedsInteractivedataGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPlanetomFeedsInteractivedataGetResponse
	FeedId          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) SetFeedId(feedId2 string) {
	this.FeedId = feedId2
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetFeedId() string {
	return this.FeedId
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.planetom.feeds.interactivedata.get"
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["feed_id"] = this.FeedId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FeedId, "feedId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPlanetomFeedsInteractivedataGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPlanetomFeedsInteractivedataGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenFeedInteractiveDataModel `json:"result,omitempty"`
	Success bool                         `json:"success,omitempty"`
}
type BaseFeedInfoModel struct {
	LiveDuration     int64    `json:"live_duration,omitempty"`
	ShareLiveCidList []string `json:"share_live_cid_list,omitempty"`
	StartTime        int64    `json:"start_time,omitempty"`
	Title            string   `json:"title,omitempty"`
}
type InteractiveInfoModel struct {
	MessageCount int64 `json:"message_count,omitempty"`
	PraiseCount  int64 `json:"praise_count,omitempty"`
	Pv           int64 `json:"pv,omitempty"`
	Uv           int64 `json:"uv,omitempty"`
}
type OpenFeedInteractiveDataModel struct {
	BaseFeedInfo    BaseFeedInfoModel    `json:"base_feed_info,omitempty"`
	InteractiveInfo InteractiveInfoModel `json:"interactive_info,omitempty"`
}
