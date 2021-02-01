package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceQrQueryRequest() *OapiSmartdeviceQrQueryRequest {
	return &OapiSmartdeviceQrQueryRequest{}
}

type OapiSmartdeviceQrQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceQrQueryResponse
	QrContent       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceQrQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceQrQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceQrQueryRequest) SetQrContent(qrContent2 string) {
	this.QrContent = qrContent2
}
func (this *OapiSmartdeviceQrQueryRequest) GetQrContent() string {
	return this.QrContent
}
func (this *OapiSmartdeviceQrQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.qr.query"
}
func (this *OapiSmartdeviceQrQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceQrQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceQrQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceQrQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceQrQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["qr_content"] = this.QrContent
	return txtParams
}
func (this *OapiSmartdeviceQrQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceQrQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceQrQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceQrQueryResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
