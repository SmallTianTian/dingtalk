package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceExternalBindRequest() *OapiSmartdeviceExternalBindRequest {
	return &OapiSmartdeviceExternalBindRequest{}
}

type OapiSmartdeviceExternalBindRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceExternalBindResponse
	DeviceBindReqVo string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceExternalBindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceExternalBindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceExternalBindRequest) SetDeviceBindReqVo(deviceBindReqVo2 string) {
	this.DeviceBindReqVo = deviceBindReqVo2
}
func (this *OapiSmartdeviceExternalBindRequest) GetDeviceBindReqVo() string {
	return this.DeviceBindReqVo
}
func (this *OapiSmartdeviceExternalBindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.external.bind"
}
func (this *OapiSmartdeviceExternalBindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceExternalBindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceExternalBindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceExternalBindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceExternalBindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_bind_req_vo"] = this.DeviceBindReqVo
	return txtParams
}
func (this *OapiSmartdeviceExternalBindRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceExternalBindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceExternalBindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceExternalBindResponse struct {
	taobao.TaobaoResponse
	Result  DeviceBindRespVo `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
