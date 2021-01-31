package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceEventPostRequest() *OapiSmartdeviceEventPostRequest {
	return &OapiSmartdeviceEventPostRequest{}
}

type OapiSmartdeviceEventPostRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceEventPostResponse
	DeviceEventVo   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceEventPostRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceEventPostRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceEventPostRequest) SetDeviceEventVo(deviceEventVo2 string) {
	this.DeviceEventVo = deviceEventVo2
}
func (this *OapiSmartdeviceEventPostRequest) GetDeviceEventVo() string {
	return this.DeviceEventVo
}
func (this *OapiSmartdeviceEventPostRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.event.post"
}
func (this *OapiSmartdeviceEventPostRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceEventPostRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceEventPostRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceEventPostRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceEventPostRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_event_vo"] = this.DeviceEventVo
	return txtParams
}
func (this *OapiSmartdeviceEventPostRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceEventPostRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceEventPostRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceEventPostResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
