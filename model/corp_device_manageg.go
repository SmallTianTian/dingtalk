package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpDeviceManageGetRequest() *CorpDeviceManageGetRequest {
	return &CorpDeviceManageGetRequest{}
}

type CorpDeviceManageGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDeviceManageGetResponse
	DeviceId        string
	DeviceServiceId int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDeviceManageGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDeviceManageGetRequest) SetDeviceId(deviceId2 string) {
	this.DeviceId = deviceId2
}
func (this *CorpDeviceManageGetRequest) GetDeviceId() string {
	return this.DeviceId
}
func (this *CorpDeviceManageGetRequest) SetDeviceServiceId(deviceServiceId2 int64) {
	this.DeviceServiceId = deviceServiceId2
}
func (this *CorpDeviceManageGetRequest) GetDeviceServiceId() int64 {
	return this.DeviceServiceId
}
func (this *CorpDeviceManageGetRequest) GetApiMethodName() string {
	return "dingtalk.corp.device.manage.get"
}
func (this *CorpDeviceManageGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDeviceManageGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDeviceManageGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDeviceManageGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDeviceManageGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDeviceManageGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_id"] = this.DeviceId
	txtParams["device_service_id"] = this.DeviceServiceId
	return txtParams
}
func (this *CorpDeviceManageGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeviceId, "deviceId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpDeviceManageGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDeviceManageGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpDeviceManageGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenDeviceDetailVo struct {
	DeviceId       string `json:"device_id,omitempty"`
	DeviceMac      string `json:"device_mac,omitempty"`
	DeviceTypeName string `json:"device_type_name,omitempty"`
	Nick           string `json:"nick,omitempty"`
	Version        string `json:"version,omitempty"`
}
