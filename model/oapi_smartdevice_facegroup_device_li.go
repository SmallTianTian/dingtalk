package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupDeviceListRequest() *OapiSmartdeviceFacegroupDeviceListRequest {
	return &OapiSmartdeviceFacegroupDeviceListRequest{}
}

type OapiSmartdeviceFacegroupDeviceListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupDeviceListResponse
	BizId           string
	Cursor          int64
	Mode            string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetMode(mode2 string) {
	this.Mode = mode2
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetMode() string {
	return this.Mode
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetType() string {
	return this.Type
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.device.list"
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["cursor"] = this.Cursor
	txtParams["mode"] = this.Mode
	txtParams["size"] = this.Size
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupDeviceListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupDeviceListResponse struct {
	taobao.TaobaoResponse
	Result  PagedList `json:"result,omitempty"`
	Success bool      `json:"success,omitempty"`
}
type DeviceDto struct {
	DeviceId int64  `json:"device_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Online   bool   `json:"online,omitempty"`
	Status   bool   `json:"status,omitempty"`
	Type     string `json:"type,omitempty"`
	Used     bool   `json:"used,omitempty"`
}
