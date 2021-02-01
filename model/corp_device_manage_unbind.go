package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpDeviceManageUnbindRequest() *CorpDeviceManageUnbindRequest {
	return &CorpDeviceManageUnbindRequest{}
}

type CorpDeviceManageUnbindRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDeviceManageUnbindResponse
	DeviceId        string
	DeviceServiceId int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDeviceManageUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDeviceManageUnbindRequest) SetDeviceId(deviceId2 string) {
	this.DeviceId = deviceId2
}
func (this *CorpDeviceManageUnbindRequest) GetDeviceId() string {
	return this.DeviceId
}
func (this *CorpDeviceManageUnbindRequest) SetDeviceServiceId(deviceServiceId2 int64) {
	this.DeviceServiceId = deviceServiceId2
}
func (this *CorpDeviceManageUnbindRequest) GetDeviceServiceId() int64 {
	return this.DeviceServiceId
}
func (this *CorpDeviceManageUnbindRequest) GetApiMethodName() string {
	return "dingtalk.corp.device.manage.unbind"
}
func (this *CorpDeviceManageUnbindRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDeviceManageUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDeviceManageUnbindRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDeviceManageUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDeviceManageUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDeviceManageUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_id"] = this.DeviceId
	txtParams["device_service_id"] = this.DeviceServiceId
	return txtParams
}
func (this *CorpDeviceManageUnbindRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeviceId, "deviceId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpDeviceManageUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDeviceManageUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpDeviceManageUnbindResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
