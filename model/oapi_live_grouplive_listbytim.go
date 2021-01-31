package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiLiveGroupliveListbytimeRequest() *OapiLiveGroupliveListbytimeRequest {
	return &OapiLiveGroupliveListbytimeRequest{}
}

type OapiLiveGroupliveListbytimeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveGroupliveListbytimeResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveGroupliveListbytimeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveGroupliveListbytimeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveGroupliveListbytimeRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiLiveGroupliveListbytimeRequest) GetReq() string {
	return this.Req
}
func (this *OapiLiveGroupliveListbytimeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.grouplive.listbytime"
}
func (this *OapiLiveGroupliveListbytimeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveGroupliveListbytimeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveGroupliveListbytimeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveGroupliveListbytimeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveGroupliveListbytimeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiLiveGroupliveListbytimeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiLiveGroupliveListbytimeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveGroupliveListbytimeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GroupLiveRecordReq struct {
	DeptId    int64 `json:"dept_id,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`
	StartTime int64 `json:"start_time,omitempty"`
}
type OapiLiveGroupliveListbytimeResponse struct {
	taobao.TaobaoResponse
	Result []GroupLiveListResult `json:"result,omitempty"`
}
type GroupLiveListResult struct {
	Duration  int64  `json:"duration,omitempty"`
	LiveUuid  string `json:"live_uuid,omitempty"`
	ShareFrom bool   `json:"share_from,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	Title     string `json:"title,omitempty"`
	Userid    string `json:"userid,omitempty"`
}
