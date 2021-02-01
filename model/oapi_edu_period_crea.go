package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduPeriodCreateRequest() *OapiEduPeriodCreateRequest {
	return &OapiEduPeriodCreateRequest{}
}

type OapiEduPeriodCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduPeriodCreateResponse
	OpenPeriod      string
	Operator        string
	SuperId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduPeriodCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduPeriodCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduPeriodCreateRequest) SetOpenPeriod(openPeriod2 string) {
	this.OpenPeriod = openPeriod2
}
func (this *OapiEduPeriodCreateRequest) GetOpenPeriod() string {
	return this.OpenPeriod
}
func (this *OapiEduPeriodCreateRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduPeriodCreateRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduPeriodCreateRequest) SetSuperId(superId2 int64) {
	this.SuperId = superId2
}
func (this *OapiEduPeriodCreateRequest) GetSuperId() int64 {
	return this.SuperId
}
func (this *OapiEduPeriodCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.period.create"
}
func (this *OapiEduPeriodCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduPeriodCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduPeriodCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduPeriodCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduPeriodCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_period"] = this.OpenPeriod
	txtParams["operator"] = this.Operator
	txtParams["super_id"] = this.SuperId
	return txtParams
}
func (this *OapiEduPeriodCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Operator, "operator"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduPeriodCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduPeriodCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenPeriod struct {
	Grades     []Grades `json:"grades,omitempty"`
	NameMode   string   `json:"name_mode,omitempty"`
	PeriodCode string   `json:"period_code,omitempty"`
	Step       string   `json:"step,omitempty"`
}
type OapiEduPeriodCreateResponse struct {
	taobao.TaobaoResponse
	Result  OpenPeriodCreateResponse `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type EduGradeDo struct {
	CampusId int64  `json:"campus_id,omitempty"`
	DeptId   int64  `json:"dept_id,omitempty"`
	Grade    int64  `json:"grade,omitempty"`
	Name     string `json:"name,omitempty"`
	SuperId  int64  `json:"super_id,omitempty"`
}
type OpenPeriodCreateResponse struct {
	DeptId int64        `json:"deptId,omitempty"`
	Grades []EduGradeDo `json:"grades,omitempty"`
}
