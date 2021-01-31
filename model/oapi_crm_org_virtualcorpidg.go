package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmOrgVirtualcorpidGetRequest() *OapiCrmOrgVirtualcorpidGetRequest {
	return &OapiCrmOrgVirtualcorpidGetRequest{}
}

type OapiCrmOrgVirtualcorpidGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmOrgVirtualcorpidGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmOrgVirtualcorpidGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.org.virtualcorpid.get"
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmOrgVirtualcorpidGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmOrgVirtualcorpidGetResponse struct {
	taobao.TaobaoResponse
	Errcode       int64  `json:"errcode,omitempty"`
	Errmsg        string `json:"errmsg,omitempty"`
	VirtualCorpid string `json:"virtual_corpid,omitempty"`
}
