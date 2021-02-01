package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRetailUserBindqueryRequest() *OapiRetailUserBindqueryRequest {
	return &OapiRetailUserBindqueryRequest{}
}

type OapiRetailUserBindqueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiRetailUserBindqueryResponse
	AssociateUnionId string
	Channel          string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiRetailUserBindqueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailUserBindqueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailUserBindqueryRequest) SetAssociateUnionId(associateUnionId2 string) {
	this.AssociateUnionId = associateUnionId2
}
func (this *OapiRetailUserBindqueryRequest) GetAssociateUnionId() string {
	return this.AssociateUnionId
}
func (this *OapiRetailUserBindqueryRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiRetailUserBindqueryRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiRetailUserBindqueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.user.bindquery"
}
func (this *OapiRetailUserBindqueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailUserBindqueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailUserBindqueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailUserBindqueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailUserBindqueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["associate_union_id"] = this.AssociateUnionId
	txtParams["channel"] = this.Channel
	return txtParams
}
func (this *OapiRetailUserBindqueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRetailUserBindqueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailUserBindqueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRetailUserBindqueryResponse struct {
	taobao.TaobaoResponse
	Result  UserBindDto `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
type ApplyUserDto struct {
	Channel      string `json:"channel,omitempty"`
	Extension    string `json:"extension,omitempty"`
	OuterId      string `json:"outer_id,omitempty"`
	OuterNick    string `json:"outer_nick,omitempty"`
	OuterSubId   string `json:"outer_sub_id,omitempty"`
	OuterSubNick string `json:"outer_sub_nick,omitempty"`
}
