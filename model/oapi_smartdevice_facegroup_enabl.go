package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupEnableRequest() *OapiSmartdeviceFacegroupEnableRequest {
	return &OapiSmartdeviceFacegroupEnableRequest{}
}

type OapiSmartdeviceFacegroupEnableRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupEnableResponse
	BizId           string
	DeviceIds       string
	Enable          bool
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFacegroupEnableRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupEnableRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupEnableRequest) SetDeviceIds(deviceIds2 string) {
	this.DeviceIds = deviceIds2
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetDeviceIds() string {
	return this.DeviceIds
}
func (this *OapiSmartdeviceFacegroupEnableRequest) SetEnable(enable2 bool) {
	this.Enable = enable2
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetEnable() bool {
	return this.Enable
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.enable"
}
func (this *OapiSmartdeviceFacegroupEnableRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupEnableRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupEnableRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	txtParams["device_ids"] = this.DeviceIds
	txtParams["enable"] = this.Enable
	return txtParams
}
func (this *OapiSmartdeviceFacegroupEnableRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizId, "bizId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupEnableRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupEnableResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
