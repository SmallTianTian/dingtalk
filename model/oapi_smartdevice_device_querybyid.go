package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceDeviceQuerybyidRequest() *OapiSmartdeviceDeviceQuerybyidRequest {
	return &OapiSmartdeviceDeviceQuerybyidRequest{}
}

type OapiSmartdeviceDeviceQuerybyidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceDeviceQuerybyidResponse
	DeviceQueryVo   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) SetDeviceQueryVo(deviceQueryVo2 string) {
	this.DeviceQueryVo = deviceQueryVo2
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetDeviceQueryVo() string {
	return this.DeviceQueryVo
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.device.querybyid"
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_query_vo"] = this.DeviceQueryVo
	return txtParams
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDeviceQuerybyidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceDeviceQuerybyidResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  DeviceDetailVO `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
