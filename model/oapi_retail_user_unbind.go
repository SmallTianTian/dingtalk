package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRetailUserUnbindRequest() *OapiRetailUserUnbindRequest {
	return &OapiRetailUserUnbindRequest{}
}

type OapiRetailUserUnbindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiRetailUserUnbindResponse
	AssociateUnionId string
	Channel          string
	OuterId          string
	SubOuterId       string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiRetailUserUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailUserUnbindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailUserUnbindRequest) SetAssociateUnionId(associateUnionId2 string) {
	this.AssociateUnionId = associateUnionId2
}
func (this *OapiRetailUserUnbindRequest) GetAssociateUnionId() string {
	return this.AssociateUnionId
}
func (this *OapiRetailUserUnbindRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiRetailUserUnbindRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiRetailUserUnbindRequest) SetOuterId(outerId2 string) {
	this.OuterId = outerId2
}
func (this *OapiRetailUserUnbindRequest) GetOuterId() string {
	return this.OuterId
}
func (this *OapiRetailUserUnbindRequest) SetSubOuterId(subOuterId2 string) {
	this.SubOuterId = subOuterId2
}
func (this *OapiRetailUserUnbindRequest) GetSubOuterId() string {
	return this.SubOuterId
}
func (this *OapiRetailUserUnbindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.user.unbind"
}
func (this *OapiRetailUserUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailUserUnbindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailUserUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailUserUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailUserUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["associate_union_id"] = this.AssociateUnionId
	txtParams["channel"] = this.Channel
	txtParams["outer_id"] = this.OuterId
	txtParams["sub_outer_id"] = this.SubOuterId
	return txtParams
}
func (this *OapiRetailUserUnbindRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRetailUserUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailUserUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRetailUserUnbindResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
