package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRetailUserBindapplyRequest() *OapiRetailUserBindapplyRequest {
	return &OapiRetailUserBindapplyRequest{}
}

type OapiRetailUserBindapplyRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRetailUserBindapplyResponse
	Channel         string
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRetailUserBindapplyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailUserBindapplyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailUserBindapplyRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiRetailUserBindapplyRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiRetailUserBindapplyRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRetailUserBindapplyRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRetailUserBindapplyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.user.bindapply"
}
func (this *OapiRetailUserBindapplyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailUserBindapplyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailUserBindapplyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailUserBindapplyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailUserBindapplyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["channel"] = this.Channel
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRetailUserBindapplyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRetailUserBindapplyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailUserBindapplyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UserBindParam struct {
	Extension    string `json:"extension,omitempty"`
	OuterId      string `json:"outer_id,omitempty"`
	OuterNick    string `json:"outer_nick,omitempty"`
	SubOuterId   string `json:"sub_outer_id,omitempty"`
	SubOuterNick string `json:"sub_outer_nick,omitempty"`
	Token        string `json:"token,omitempty"`
}
type OapiRetailUserBindapplyResponse struct {
	taobao.TaobaoResponse
	Result  UserBindDto `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
type UserBindDto struct {
	AssociateUnionId string `json:"associate_union_id,omitempty"`
	Extension        string `json:"extension,omitempty"`
}
