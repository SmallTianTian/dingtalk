package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingpayVirtualaccountQueryRequest() *OapiDingpayVirtualaccountQueryRequest {
	return &OapiDingpayVirtualaccountQueryRequest{}
}

type OapiDingpayVirtualaccountQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingpayVirtualaccountQueryResponse
	Extension       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingpayVirtualaccountQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayVirtualaccountQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayVirtualaccountQueryRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayVirtualaccountQueryRequest) SetExtensionString(extension2 string) {
	this.Extension = extension2
}
func (this *OapiDingpayVirtualaccountQueryRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiDingpayVirtualaccountQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.virtualaccount.query"
}
func (this *OapiDingpayVirtualaccountQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayVirtualaccountQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayVirtualaccountQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayVirtualaccountQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayVirtualaccountQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["extension"] = this.Extension
	return txtParams
}
func (this *OapiDingpayVirtualaccountQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingpayVirtualaccountQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayVirtualaccountQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayVirtualaccountQueryResponse struct {
	taobao.TaobaoResponse
	Result  AccountQueryOpenResponse `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type DingPayAccountOpenBo struct {
	AnonymousAlipayUid string   `json:"anonymous_alipay_uid,omitempty"`
	CorpId             string   `json:"corp_id,omitempty"`
	Extension          string   `json:"extension,omitempty"`
	RealAlipayUids     []string `json:"real_alipay_uids,omitempty"`
	RealUsedAlipayUid  string   `json:"real_used_alipay_uid,omitempty"`
}
type AccountQueryOpenResponse struct {
	AccountOpenBo DingPayAccountOpenBo `json:"account_open_bo,omitempty"`
}
