package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiPbpEventDeleteRequest() *OapiPbpEventDeleteRequest {
	return &OapiPbpEventDeleteRequest{}
}

type OapiPbpEventDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpEventDeleteResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpEventDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpEventDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpEventDeleteRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiPbpEventDeleteRequest) GetParam() string {
	return this.Param
}
func (this *OapiPbpEventDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.event.delete"
}
func (this *OapiPbpEventDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpEventDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpEventDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpEventDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpEventDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiPbpEventDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiPbpEventDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpEventDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UserEventOapiVo struct {
	BizInstId string `json:"biz_inst_id,omitempty"`
	EventId   string `json:"event_id,omitempty"`
	EventName string `json:"event_name,omitempty"`
	Userid    string `json:"userid,omitempty"`
}
type UserEventOapiRequestVo struct {
	BizCode       string            `json:"biz_code,omitempty"`
	UserEventList []UserEventOapiVo `json:"user_event_list,omitempty"`
}
type OapiPbpEventDeleteResponse struct {
	taobao.TaobaoResponse
}
