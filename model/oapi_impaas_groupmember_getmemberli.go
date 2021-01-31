package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasGroupmemberGetmemberlistRequest() *OapiImpaasGroupmemberGetmemberlistRequest {
	return &OapiImpaasGroupmemberGetmemberlistRequest{}
}

type OapiImpaasGroupmemberGetmemberlistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupmemberGetmemberlistResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.groupmember.getmemberlist"
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupmemberGetmemberlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetGroupMemberListRequest struct {
	Chatid string `json:"chatid,omitempty"`
	Length int64  `json:"length,omitempty"`
	Offset int64  `json:"offset,omitempty"`
}
type OapiImpaasGroupmemberGetmemberlistResponse struct {
	taobao.TaobaoResponse
	Errcode    int64             `json:"errcode,omitempty"`
	Errmsg     string            `json:"errmsg,omitempty"`
	MemberList []GroupMemberInfo `json:"member_list,omitempty"`
}
type GroupMemberInfo struct {
	Ext  string `json:"ext,omitempty"`
	Id   string `json:"id,omitempty"`
	Nick string `json:"nick,omitempty"`
	Role int64  `json:"role,omitempty"`
	Type string `json:"type,omitempty"`
}
