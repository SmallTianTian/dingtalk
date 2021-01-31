package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceDeviceCheckOutRequest() *OapiRhinoMosSpaceDeviceCheckOutRequest {
	return &OapiRhinoMosSpaceDeviceCheckOutRequest{}
}

type OapiRhinoMosSpaceDeviceCheckOutRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceDeviceCheckOutResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.device.check.out"
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceDeviceCheckOutRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosSpaceDeviceCheckOutResponse struct {
	taobao.TaobaoResponse
	Model bool `json:"model,omitempty"`
}
