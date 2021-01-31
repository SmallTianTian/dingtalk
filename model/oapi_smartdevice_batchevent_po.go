package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceBatcheventPostRequest() *OapiSmartdeviceBatcheventPostRequest {
	return &OapiSmartdeviceBatcheventPostRequest{}
}

type OapiSmartdeviceBatcheventPostRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceBatcheventPostResponse
	DeviceEventVos  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceBatcheventPostRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceBatcheventPostRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceBatcheventPostRequest) SetDeviceEventVos(deviceEventVos2 string) {
	this.DeviceEventVos = deviceEventVos2
}
func (this *OapiSmartdeviceBatcheventPostRequest) GetDeviceEventVos() string {
	return this.DeviceEventVos
}
func (this *OapiSmartdeviceBatcheventPostRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.batchevent.post"
}
func (this *OapiSmartdeviceBatcheventPostRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceBatcheventPostRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceBatcheventPostRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceBatcheventPostRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceBatcheventPostRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["device_event_vos"] = this.DeviceEventVos
	return txtParams
}
func (this *OapiSmartdeviceBatcheventPostRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.DeviceEventVos, 999, "deviceEventVos"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceBatcheventPostRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceBatcheventPostRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeviceEventVo struct {
	Data  string `json:"data,omitempty"`
	Dn    string `json:"dn,omitempty"`
	Pk    string `json:"pk,omitempty"`
	Topic string `json:MessageFields.DATA_TOPI,omitemptyC`
}
type OapiSmartdeviceBatcheventPostResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
