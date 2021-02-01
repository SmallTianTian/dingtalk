package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasGroupQueryRequest() *OapiImpaasGroupQueryRequest {
	return &OapiImpaasGroupQueryRequest{}
}

type OapiImpaasGroupQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupQueryResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupQueryRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasGroupQueryRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasGroupQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.group.query"
}
func (this *OapiImpaasGroupQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasGroupQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasGroupQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetGroupInfoRequest struct {
	Chatid string `json:"chatid,omitempty"`
}
type OapiImpaasGroupQueryResponse struct {
	taobao.TaobaoResponse
	Result GroupInfo `json:"result,omitempty"`
}
type GroupInfo struct {
	Chatid      string              `json:"chatid,omitempty"`
	Creater     BaseGroupMemberInfo `json:"creater,omitempty"`
	GroupName   string              `json:"group_name,omitempty"`
	MemberCount int64               `json:"member_count,omitempty"`
	MemberLimit int64               `json:"member_limit,omitempty"`
	Type        int64               `json:"type,omitempty"`
}
