package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpDeviceManageHasbinddeviceRequest() *CorpDeviceManageHasbinddeviceRequest {
	return &CorpDeviceManageHasbinddeviceRequest{}
}

type CorpDeviceManageHasbinddeviceRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDeviceManageHasbinddeviceResponse
	DeviceServiceId int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDeviceManageHasbinddeviceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDeviceManageHasbinddeviceRequest) SetDeviceServiceId(deviceServiceId2 int64) {
	this.DeviceServiceId = deviceServiceId2
}
func (this *CorpDeviceManageHasbinddeviceRequest) GetDeviceServiceId() int64 {
	return this.DeviceServiceId
}
func (this *CorpDeviceManageHasbinddeviceRequest) GetApiMethodName() string {
	return "dingtalk.corp.device.manage.hasbinddevice"
}
func (this *CorpDeviceManageHasbinddeviceRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDeviceManageHasbinddeviceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDeviceManageHasbinddeviceRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDeviceManageHasbinddeviceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDeviceManageHasbinddeviceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDeviceManageHasbinddeviceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_service_id"] = this.DeviceServiceId
	return txtParams
}
func (this *CorpDeviceManageHasbinddeviceRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeviceServiceId, "deviceServiceId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpDeviceManageHasbinddeviceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDeviceManageHasbinddeviceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpDeviceManageHasbinddeviceResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
