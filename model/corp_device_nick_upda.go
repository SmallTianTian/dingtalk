package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpDeviceNickUpdateRequest() *CorpDeviceNickUpdateRequest {
	return &CorpDeviceNickUpdateRequest{}
}

type CorpDeviceNickUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDeviceNickUpdateResponse
	DeviceId        string
	DeviceServiceId int64
	NewNick         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDeviceNickUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDeviceNickUpdateRequest) SetDeviceId(deviceId2 string) {
	this.DeviceId = deviceId2
}
func (this *CorpDeviceNickUpdateRequest) GetDeviceId() string {
	return this.DeviceId
}
func (this *CorpDeviceNickUpdateRequest) SetDeviceServiceId(deviceServiceId2 int64) {
	this.DeviceServiceId = deviceServiceId2
}
func (this *CorpDeviceNickUpdateRequest) GetDeviceServiceId() int64 {
	return this.DeviceServiceId
}
func (this *CorpDeviceNickUpdateRequest) SetNewNick(newNick2 string) {
	this.NewNick = newNick2
}
func (this *CorpDeviceNickUpdateRequest) GetNewNick() string {
	return this.NewNick
}
func (this *CorpDeviceNickUpdateRequest) GetApiMethodName() string {
	return "dingtalk.corp.device.nick.update"
}
func (this *CorpDeviceNickUpdateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDeviceNickUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDeviceNickUpdateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDeviceNickUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDeviceNickUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDeviceNickUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_id"] = this.DeviceId
	txtParams["device_service_id"] = this.DeviceServiceId
	txtParams["new_nick"] = this.NewNick
	return txtParams
}
func (this *CorpDeviceNickUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeviceId, "deviceId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpDeviceNickUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDeviceNickUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpDeviceNickUpdateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
