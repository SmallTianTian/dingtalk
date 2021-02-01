package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCallbackFailrecordListRequest() *OapiCallbackFailrecordListRequest {
	return &OapiCallbackFailrecordListRequest{}
}

type OapiCallbackFailrecordListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallbackFailrecordListResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallbackFailrecordListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallbackFailrecordListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallbackFailrecordListRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiCallbackFailrecordListRequest) GetReq() string {
	return this.Req
}
func (this *OapiCallbackFailrecordListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.callback.failrecord.list"
}
func (this *OapiCallbackFailrecordListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallbackFailrecordListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallbackFailrecordListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallbackFailrecordListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallbackFailrecordListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiCallbackFailrecordListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCallbackFailrecordListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallbackFailrecordListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Req struct {
	BeginTime int64    `json:"begin_time,omitempty"`
	Confirm   bool     `json:"confirm,omitempty"`
	Cursor    int64    `json:"cursor,omitempty"`
	EndTime   int64    `json:"end_time,omitempty"`
	PageSize  int64    `json:"page_size,omitempty"`
	Status    int64    `json:"status,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}
type OapiCallbackFailrecordListResponse struct {
	taobao.TaobaoResponse

	FailedList []FailedList `json:"failed_list,omitempty"`
	HasMore    bool         `json:"has_more,omitempty"`
}
type FailedList struct {
	CallBackData string `json:"call_back_data,omitempty"`
	CallBackTag  string `json:"call_back_tag,omitempty"`
	Corpid       string `json:"corpid,omitempty"`
	EventTime    int64  `json:"event_time,omitempty"`
	Id           int64  `json:"id,omitempty"`
}
