package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceDeviceQueryRequest() *OapiSmartdeviceDeviceQueryRequest {
	return &OapiSmartdeviceDeviceQueryRequest{}
}

type OapiSmartdeviceDeviceQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceDeviceQueryResponse
	DeviceQueryVo   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceDeviceQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDeviceQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDeviceQueryRequest) SetDeviceQueryVo(deviceQueryVo2 string) {
	this.DeviceQueryVo = deviceQueryVo2
}
func (this *OapiSmartdeviceDeviceQueryRequest) GetDeviceQueryVo() string {
	return this.DeviceQueryVo
}
func (this *OapiSmartdeviceDeviceQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.device.query"
}
func (this *OapiSmartdeviceDeviceQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDeviceQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDeviceQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDeviceQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDeviceQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_query_vo"] = this.DeviceQueryVo
	return txtParams
}
func (this *OapiSmartdeviceDeviceQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceDeviceQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDeviceQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeviceQueryVo struct {
	DeviceId   string `json:"device_id,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	Pk         string `json:"pk,omitempty"`
}
type OapiSmartdeviceDeviceQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  DeviceDetailVO `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type DeviceDetailVO struct {
	CorpId     string `json:"corp_id,omitempty"`
	DeviceId   string `json:"device_id,omitempty"`
	DeviceMac  string `json:"device_mac,omitempty"`
	DeviceName string `json:"device_name,omitempty"`
	Ext        string `json:"ext,omitempty"`
	Nick       string `json:"nick,omitempty"`
	Pk         string `json:"pk,omitempty"`
	Sn         string `json:"sn,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
