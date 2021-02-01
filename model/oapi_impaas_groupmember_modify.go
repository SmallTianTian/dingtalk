package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasGroupmemberModifyRequest() *OapiImpaasGroupmemberModifyRequest {
	return &OapiImpaasGroupmemberModifyRequest{}
}

type OapiImpaasGroupmemberModifyRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupmemberModifyResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupmemberModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupmemberModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupmemberModifyRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasGroupmemberModifyRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasGroupmemberModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.groupmember.modify"
}
func (this *OapiImpaasGroupmemberModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupmemberModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupmemberModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupmemberModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupmemberModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasGroupmemberModifyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasGroupmemberModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupmemberModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GroupMemberListModifyRequest struct {
	Channel    string                `json:"channel,omitempty"`
	Chatid     string                `json:"chatid,omitempty"`
	MemberList []BaseGroupMemberInfo `json:"member_list,omitempty"`
	ModifyType string                `json:"modify_type,omitempty"`
}
type OapiImpaasGroupmemberModifyResponse struct {
	taobao.TaobaoResponse
}
