package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartworkHrmEmployeeAddpreentryRequest() *OapiSmartworkHrmEmployeeAddpreentryRequest {
	return &OapiSmartworkHrmEmployeeAddpreentryRequest{}
}

type OapiSmartworkHrmEmployeeAddpreentryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeAddpreentryResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetParam() string {
	return this.Param
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.addpreentry"
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeAddpreentryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PreEntryEmployeeAddParam struct {
	ExtendInfo   string    `json:"extend_info,omitempty"`
	Mobile       string    `json:"mobile,omitempty"`
	Name         string    `json:"name,omitempty"`
	OpUserid     string    `json:"op_userid,omitempty"`
	PreEntryTime time.Time `json:"pre_entry_time,omitempty"`
}
type OapiSmartworkHrmEmployeeAddpreentryResponse struct {
	taobao.TaobaoResponse
	Success bool   `json:"success,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
