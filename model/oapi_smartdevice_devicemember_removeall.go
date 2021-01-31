package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceDevicememberRemoveallRequest() *OapiSmartdeviceDevicememberRemoveallRequest {
	return &OapiSmartdeviceDevicememberRemoveallRequest{}
}

type OapiSmartdeviceDevicememberRemoveallRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceDevicememberRemoveallResponse
	DeviceId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) SetDeviceId(deviceId2 int64) {
	this.DeviceId = deviceId2
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetDeviceId() int64 {
	return this.DeviceId
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.devicemember.removeall"
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_id"] = this.DeviceId
	return txtParams
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeviceId, "deviceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDevicememberRemoveallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceDevicememberRemoveallResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
}
