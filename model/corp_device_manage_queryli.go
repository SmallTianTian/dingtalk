package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpDeviceManageQuerylistRequest() *CorpDeviceManageQuerylistRequest {
	return &CorpDeviceManageQuerylistRequest{}
}

type CorpDeviceManageQuerylistRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDeviceManageQuerylistResponse
	Cursor          int64
	DeviceServiceId int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDeviceManageQuerylistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDeviceManageQuerylistRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *CorpDeviceManageQuerylistRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *CorpDeviceManageQuerylistRequest) SetDeviceServiceId(deviceServiceId2 int64) {
	this.DeviceServiceId = deviceServiceId2
}
func (this *CorpDeviceManageQuerylistRequest) GetDeviceServiceId() int64 {
	return this.DeviceServiceId
}
func (this *CorpDeviceManageQuerylistRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpDeviceManageQuerylistRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpDeviceManageQuerylistRequest) GetApiMethodName() string {
	return "dingtalk.corp.device.manage.querylist"
}
func (this *CorpDeviceManageQuerylistRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDeviceManageQuerylistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDeviceManageQuerylistRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDeviceManageQuerylistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDeviceManageQuerylistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDeviceManageQuerylistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["device_service_id"] = this.DeviceServiceId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpDeviceManageQuerylistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *CorpDeviceManageQuerylistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDeviceManageQuerylistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpDeviceManageQuerylistResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OpenDeviceVo struct {
	Avatar         string `json:"avatar,omitempty"`
	DeviceId       string `json:"device_id,omitempty"`
	DeviceMac      string `json:"device_mac,omitempty"`
	DeviceSn       string `json:"device_sn,omitempty"`
	DeviceTypeName string `json:"device_type_name,omitempty"`
	Nick           string `json:"nick,omitempty"`
	OnLineStatus   int64  `json:"on_line_status,omitempty"`
}
type PageVo struct {
	List       []OpenDeviceVo `json:"list,omitempty"`
	NextCursor int64          `json:"next_cursor,omitempty"`
}
