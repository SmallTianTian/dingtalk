package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceDeviceCheckInRequest() *OapiRhinoMosSpaceDeviceCheckInRequest {
	return &OapiRhinoMosSpaceDeviceCheckInRequest{}
}

type OapiRhinoMosSpaceDeviceCheckInRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceDeviceCheckInResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.device.check.in"
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceDeviceCheckInRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProdWorkstationDeviceBatchCheckReq struct {
	DeviceIds           []int64 `json:"device_ids,omitempty"`
	ProdWorkstationCode string  `json:"prod_workstation_code,omitempty"`
	TenantId            string  `json:"tenant_id,omitempty"`
	Userid              string  `json:"userid,omitempty"`
}
type OapiRhinoMosSpaceDeviceCheckInResponse struct {
	taobao.TaobaoResponse
	Model bool `json:"model,omitempty"`
}
