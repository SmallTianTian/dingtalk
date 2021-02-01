package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupDeviceUpdateRequest() *OapiSmartdeviceFacegroupDeviceUpdateRequest {
	return &OapiSmartdeviceFacegroupDeviceUpdateRequest{}
}

type OapiSmartdeviceFacegroupDeviceUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupDeviceUpdateResponse
	AddDeviceIds    string
	BizId           string
	DelDeviceIds    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) SetAddDeviceIds(addDeviceIds2 string) {
	this.AddDeviceIds = addDeviceIds2
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetAddDeviceIds() string {
	return this.AddDeviceIds
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) SetDelDeviceIds(delDeviceIds2 string) {
	this.DelDeviceIds = delDeviceIds2
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetDelDeviceIds() string {
	return this.DelDeviceIds
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.device.update"
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["add_device_ids"] = this.AddDeviceIds
	txtParams["biz_id"] = this.BizId
	txtParams["del_device_ids"] = this.DelDeviceIds
	return txtParams
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AddDeviceIds, 20, "addDeviceIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupDeviceUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupDeviceUpdateResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
