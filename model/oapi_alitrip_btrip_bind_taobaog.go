package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripBindTaobaoGetRequest() *OapiAlitripBtripBindTaobaoGetRequest {
	return &OapiAlitripBtripBindTaobaoGetRequest{}
}

type OapiAlitripBtripBindTaobaoGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripBindTaobaoGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripBindTaobaoGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.bind.taobao.get"
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripBindTaobaoGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type UserBindInfoRq struct {
	Corpid string `json:"corpid,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type OapiAlitripBtripBindTaobaoGetResponse struct {
	taobao.TaobaoResponse
	Result  UserTaoBaoInfo `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type UserTaoBaoInfo struct {
	AlipayInfo string `json:"alipay_info,omitempty"`
	TaobaoInfo string `json:"taobao_info,omitempty"`
}
