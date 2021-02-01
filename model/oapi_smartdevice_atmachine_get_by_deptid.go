package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartdeviceAtmachineGetByDeptidRequest() *OapiSmartdeviceAtmachineGetByDeptidRequest {
	return &OapiSmartdeviceAtmachineGetByDeptidRequest{}
}

type OapiSmartdeviceAtmachineGetByDeptidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceAtmachineGetByDeptidResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.atmachine.get_by_deptid"
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceAtmachineGetByDeptidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type MachineInfoRequestVo struct {
	Deptid int64 `json:"deptid,omitempty"`
	Offset int64 `json:"offset,omitempty"`
	Size   int64 `json:"size,omitempty"`
}
type OapiSmartdeviceAtmachineGetByDeptidResponse struct {
	taobao.TaobaoResponse
	Result MachineInfoResultVo `json:"result,omitempty"`
}
type MachineVo struct {
	DeviceName  string `json:"device_name,omitempty"`
	Deviceid    string `json:"deviceid,omitempty"`
	Devid       int64  `json:"devid,omitempty"`
	ProductName string `json:"product_name,omitempty"`
}
type MachineInfoResultVo struct {
	HasMore     bool        `json:"has_more,omitempty"`
	MachineList []MachineVo `json:"machine_list,omitempty"`
}
