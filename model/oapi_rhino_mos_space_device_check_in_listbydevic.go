package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceDeviceCheckInListbydeviceRequest() *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest {
	return &OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest{}
}

type OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceDeviceCheckInListbydeviceResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.device.check.in.listbydevice"
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceDeviceCheckInListbydeviceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type WorkstationDeviceListReq struct {
	DeviceIds []int64 `json:"device_ids,omitempty"`
	TenantId  string  `json:"tenant_id,omitempty"`
	Userid    string  `json:"userid,omitempty"`
}
type OapiRhinoMosSpaceDeviceCheckInListbydeviceResponse struct {
	taobao.TaobaoResponse
	Errcode int64                           `json:"errcode,omitempty"`
	Errmsg  string                          `json:"errmsg,omitempty"`
	Model   []ProdWorkstationDeviceCheckDto `json:"model,omitempty"`
}
