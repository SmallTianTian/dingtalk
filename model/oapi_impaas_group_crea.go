package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasGroupCreateRequest() *OapiImpaasGroupCreateRequest {
	return &OapiImpaasGroupCreateRequest{}
}

type OapiImpaasGroupCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasGroupCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasGroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.group.create"
}
func (this *OapiImpaasGroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasGroupCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasGroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BaseGroupMemberInfo struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
type CreateGroupRequest struct {
	Channel         string                `json:"channel,omitempty"`
	Creater         BaseGroupMemberInfo   `json:"creater,omitempty"`
	EntranceId      int64                 `json:"entrance_id,omitempty"`
	Extension       string                `json:"extension,omitempty"`
	MemberList      []BaseGroupMemberInfo `json:"member_list,omitempty"`
	Name            string                `json:"name,omitempty"`
	ShowHistoryType int64                 `json:"show_history_type,omitempty"`
	Type            int64                 `json:"type,omitempty"`
	Uuid            string                `json:"uuid,omitempty"`
}
type OapiImpaasGroupCreateResponse struct {
	taobao.TaobaoResponse
	Chatid  string `json:"chatid,omitempty"`
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
