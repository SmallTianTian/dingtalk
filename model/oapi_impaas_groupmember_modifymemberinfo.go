package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasGroupmemberModifymemberinfoRequest() *OapiImpaasGroupmemberModifymemberinfoRequest {
	return &OapiImpaasGroupmemberModifymemberinfoRequest{}
}

type OapiImpaasGroupmemberModifymemberinfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupmemberModifymemberinfoResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.groupmember.modifymemberinfo"
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupmemberModifymemberinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ModifyGroupMemberInfoRequest struct {
	Chatid     string          `json:"chatid,omitempty"`
	MemberInfo GroupMemberInfo `json:"member_info,omitempty"`
}
type OapiImpaasGroupmemberModifymemberinfoResponse struct {
	taobao.TaobaoResponse
}
