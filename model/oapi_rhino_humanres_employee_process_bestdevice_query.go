package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest() *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest {
	return &OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest{}
}

type OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                         OapiRhinoHumanresEmployeeProcessBestdeviceQueryResponse
	AvailableDeviceModels        string
	EmployeeProcessCapacityUnits string
	TenantId                     string
	TopHttpMethod                string
	TopResponseType              string
	Userid                       string
}

func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) SetAvailableDeviceModels(availableDeviceModels2 string) {
	this.AvailableDeviceModels = availableDeviceModels2
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetAvailableDeviceModels() string {
	return this.AvailableDeviceModels
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) SetEmployeeProcessCapacityUnits(employeeProcessCapacityUnits2 string) {
	this.EmployeeProcessCapacityUnits = employeeProcessCapacityUnits2
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetEmployeeProcessCapacityUnits() string {
	return this.EmployeeProcessCapacityUnits
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.humanres.employee.process.bestdevice.query"
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["available_device_models"] = this.AvailableDeviceModels
	txtParams["employee_process_capacity_units"] = this.EmployeeProcessCapacityUnits
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.AvailableDeviceModels, 20, "availableDeviceModels"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoHumanresEmployeeProcessBestdeviceQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DeviceModelDto struct {
	DeviceModelId      string `json:"device_model_id,omitempty"`
	DeviceModelVersion string `json:"device_model_version,omitempty"`
}
type EmployeeProcessCapacityUnitReq struct {
	ProcessCapacityId int64  `json:"process_capacity_id,omitempty"`
	WorkNo            string `json:"work_no,omitempty"`
}
type OapiRhinoHumanresEmployeeProcessBestdeviceQueryResponse struct {
	taobao.TaobaoResponse
	Errcode         int64                              `json:"errcode,omitempty"`
	Errmsg          string                             `json:"errmsg,omitempty"`
	ExternalMsgInfo string                             `json:"external_msg_info,omitempty"`
	Hsfcode         int64                              `json:"hsfcode,omitempty"`
	Model           []EmployeeProcessDeviceCapacityDto `json:"model,omitempty"`
	Success         bool                               `json:"success,omitempty"`
}
type EmployeeProcessDeviceCapacityDto struct {
	DeviceModelId      string `json:"device_model_id,omitempty"`
	DeviceModelVersion string `json:"device_model_version,omitempty"`
	ProcessCapacityId  int64  `json:"process_capacity_id,omitempty"`
	ProductionQuantity int64  `json:"production_quantity,omitempty"`
	TenantId           string `json:"tenant_id,omitempty"`
	WorkNo             string `json:"work_no,omitempty"`
}
