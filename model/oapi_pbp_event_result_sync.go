package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiPbpEventResultSyncRequest() *OapiPbpEventResultSyncRequest {
	return &OapiPbpEventResultSyncRequest{}
}

type OapiPbpEventResultSyncRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpEventResultSyncResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpEventResultSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpEventResultSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpEventResultSyncRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiPbpEventResultSyncRequest) GetParam() string {
	return this.Param
}
func (this *OapiPbpEventResultSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.event.result.sync"
}
func (this *OapiPbpEventResultSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpEventResultSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpEventResultSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpEventResultSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpEventResultSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiPbpEventResultSyncRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiPbpEventResultSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpEventResultSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PositionOapiVo struct {
	PositionId   string `json:"position_id,omitempty"`
	PositionType int64  `json:"position_type,omitempty"`
}
type UserEventResultOapiRequestVo struct {
	BizCode       string         `json:"biz_code,omitempty"`
	BizInstId     string         `json:"biz_inst_id,omitempty"`
	EventId       string         `json:"event_id,omitempty"`
	InvalidEvent  bool           `json:"invalid_event,omitempty"`
	PunchPosition PositionOapiVo `json:"punch_position,omitempty"`
	Result        int64          `json:"result,omitempty"`
	Userid        string         `json:"userid,omitempty"`
}
type OapiPbpEventResultSyncResponse struct {
	taobao.TaobaoResponse
}
