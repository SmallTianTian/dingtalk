package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmAuthGroupListRequest() *OapiCrmAuthGroupListRequest {
	return &OapiCrmAuthGroupListRequest{}
}

type OapiCrmAuthGroupListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmAuthGroupListResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmAuthGroupListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmAuthGroupListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmAuthGroupListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.auth.group.list"
}
func (this *OapiCrmAuthGroupListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmAuthGroupListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmAuthGroupListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmAuthGroupListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmAuthGroupListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCrmAuthGroupListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmAuthGroupListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmAuthGroupListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmAuthGroupListResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  []RoleGroupVo `json:"result,omitempty"`
}
type RoleGroupVo struct {
	Name   string `json:"name,omitempty"`
	RoleId int64  `json:"role_id,omitempty"`
}
