package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosSpaceDeviceCheckInListRequest() *OapiRhinoMosSpaceDeviceCheckInListRequest {
	return &OapiRhinoMosSpaceDeviceCheckInListRequest{}
}

type OapiRhinoMosSpaceDeviceCheckInListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosSpaceDeviceCheckInListResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetRequest() string {
	return this.Request
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.space.device.check.in.list"
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosSpaceDeviceCheckInListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProdWorkstationDeviceListReq struct {
	ProdWorkstationCode string `json:"prod_workstation_code,omitempty"`
	TenantId            string `json:"tenant_id,omitempty"`
	Userid              string `json:"userid,omitempty"`
}
type OapiRhinoMosSpaceDeviceCheckInListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                           `json:"errcode,omitempty"`
	Errmsg  string                          `json:"errmsg,omitempty"`
	Model   []ProdWorkstationDeviceCheckDto `json:"model,omitempty"`
}
type ProdWorkstationDeviceCheckDto struct {
	CheckInTime         time.Time `json:"check_in_time,omitempty"`
	CheckOutTime        time.Time `json:"check_out_time,omitempty"`
	CheckStatus         string    `json:"check_status,omitempty"`
	DeviceId            int64     `json:"device_id,omitempty"`
	ProdWorkstationCode string    `json:"prod_workstation_code,omitempty"`
	TenantId            string    `json:"tenant_id,omitempty"`
}
