package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasGroupModifyRequest() *OapiImpaasGroupModifyRequest {
	return &OapiImpaasGroupModifyRequest{}
}

type OapiImpaasGroupModifyRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupModifyResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupModifyRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasGroupModifyRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasGroupModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.group.modify"
}
func (this *OapiImpaasGroupModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasGroupModifyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasGroupModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GroupInfoModifyRequest struct {
	Chatid          string              `json:"chatid,omitempty"`
	GroupName       string              `json:"group_name,omitempty"`
	GroupOwner      BaseGroupMemberInfo `json:"group_owner,omitempty"`
	ShowHistoryType int64               `json:"show_history_type,omitempty"`
}
type OapiImpaasGroupModifyResponse struct {
	taobao.TaobaoResponse
}
