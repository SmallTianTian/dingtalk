package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiLiveGroupliveDetailGetRequest() *OapiLiveGroupliveDetailGetRequest {
	return &OapiLiveGroupliveDetailGetRequest{}
}

type OapiLiveGroupliveDetailGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveGroupliveDetailGetResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveGroupliveDetailGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveGroupliveDetailGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveGroupliveDetailGetRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiLiveGroupliveDetailGetRequest) GetReq() string {
	return this.Req
}
func (this *OapiLiveGroupliveDetailGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.grouplive.detail.get"
}
func (this *OapiLiveGroupliveDetailGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveGroupliveDetailGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveGroupliveDetailGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveGroupliveDetailGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveGroupliveDetailGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiLiveGroupliveDetailGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiLiveGroupliveDetailGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveGroupliveDetailGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GroupLiveStatisticsReq struct {
	DeptId   int64  `json:"dept_id,omitempty"`
	LiveUuid string `json:"live_uuid,omitempty"`
}
type OapiLiveGroupliveDetailGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64               `json:"errcode,omitempty"`
	Errmsg  string              `json:"errmsg,omitempty"`
	Result  GroupLiveStatistics `json:"result,omitempty"`
}
type Groupliveviewers struct {
	PlayDuration       int64  `json:"play_duration,omitempty"`
	PlayRecordDuration int64  `json:"play_record_duration,omitempty"`
	Userid             string `json:"userid,omitempty"`
}
type GroupLiveStatistics struct {
	GroupLiveViewers []Groupliveviewers `json:"group_live_viewers,omitempty"`
	LiveShareUrl     string             `json:"live_share_url,omitempty"`
	MessageCount     int64              `json:"message_count,omitempty"`
	PraiseCount      int64              `json:"praise_count,omitempty"`
	Pv               int64              `json:"pv,omitempty"`
	UnviewedCount    int64              `json:"unviewed_count,omitempty"`
	ViewerCount      int64              `json:"viewer_count,omitempty"`
}
