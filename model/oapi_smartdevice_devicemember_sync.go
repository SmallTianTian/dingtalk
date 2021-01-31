package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceDevicememberSyncRequest() *OapiSmartdeviceDevicememberSyncRequest {
	return &OapiSmartdeviceDevicememberSyncRequest{}
}

type OapiSmartdeviceDevicememberSyncRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceDevicememberSyncResponse
	AddUserids      string
	DelUserids      string
	DeviceId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceDevicememberSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceDevicememberSyncRequest) SetAddUserids(addUserids2 string) {
	this.AddUserids = addUserids2
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetAddUserids() string {
	return this.AddUserids
}
func (this *OapiSmartdeviceDevicememberSyncRequest) SetDelUserids(delUserids2 string) {
	this.DelUserids = delUserids2
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetDelUserids() string {
	return this.DelUserids
}
func (this *OapiSmartdeviceDevicememberSyncRequest) SetDeviceId(deviceId2 int64) {
	this.DeviceId = deviceId2
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetDeviceId() int64 {
	return this.DeviceId
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.devicemember.sync"
}
func (this *OapiSmartdeviceDevicememberSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceDevicememberSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceDevicememberSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["add_userids"] = this.AddUserids
	txtParams["del_userids"] = this.DelUserids
	txtParams["device_id"] = this.DeviceId
	return txtParams
}
func (this *OapiSmartdeviceDevicememberSyncRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AddUserids, 200, "addUserids"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceDevicememberSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceDevicememberSyncResponse struct {
	taobao.TaobaoResponse
	Result []string `json:"result,omitempty"`
}
