package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceDevicememberListRequest() *OapiSmartdeviceDevicememberListRequest {
	return &OapiSmartdeviceDevicememberListRequest{}
}

type OapiSmartdeviceDevicememberListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceDevicememberListResponse
	Cursor          int64
	DeviceId        int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceDevicememberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDevicememberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDevicememberListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiSmartdeviceDevicememberListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiSmartdeviceDevicememberListRequest) SetDeviceId(deviceId2 int64) {
	this.DeviceId = deviceId2
}
func (this *OapiSmartdeviceDevicememberListRequest) GetDeviceId() int64 {
	return this.DeviceId
}
func (this *OapiSmartdeviceDevicememberListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartdeviceDevicememberListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartdeviceDevicememberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.devicemember.list"
}
func (this *OapiSmartdeviceDevicememberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDevicememberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDevicememberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDevicememberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDevicememberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["device_id"] = this.DeviceId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiSmartdeviceDevicememberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceDevicememberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDevicememberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceDevicememberListResponse struct {
	taobao.TaobaoResponse
	Result PagedList `json:"result,omitempty"`
}
type PagedList struct {
	Cursor  int64    `json:"cursor,omitempty"`
	HasMore bool     `json:"has_more,omitempty"`
	Items   []string `json:"items,omitempty"`
}
