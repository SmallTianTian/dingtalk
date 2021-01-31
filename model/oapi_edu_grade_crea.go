package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGradeCreateRequest() *OapiEduGradeCreateRequest {
	return &OapiEduGradeCreateRequest{}
}

type OapiEduGradeCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduGradeCreateResponse
	OpenGrade       string
	Operator        string
	SuperId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduGradeCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGradeCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGradeCreateRequest) SetOpenGrade(openGrade2 string) {
	this.OpenGrade = openGrade2
}
func (this *OapiEduGradeCreateRequest) GetOpenGrade() string {
	return this.OpenGrade
}
func (this *OapiEduGradeCreateRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduGradeCreateRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduGradeCreateRequest) SetSuperId(superId2 int64) {
	this.SuperId = superId2
}
func (this *OapiEduGradeCreateRequest) GetSuperId() int64 {
	return this.SuperId
}
func (this *OapiEduGradeCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.grade.create"
}
func (this *OapiEduGradeCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGradeCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGradeCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGradeCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGradeCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_grade"] = this.OpenGrade
	txtParams["operator"] = this.Operator
	txtParams["super_id"] = this.SuperId
	return txtParams
}
func (this *OapiEduGradeCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Operator, "operator"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGradeCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGradeCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenGrade struct {
	Classes   int64  `json:"classes,omitempty"`
	Grade     string `json:"grade,omitempty"`
	Name      string `json:"name,omitempty"`
	StartYear string `json:"start_year,omitempty"`
}
type OapiEduGradeCreateResponse struct {
	taobao.TaobaoResponse
	Result  OpenGradeCreateResponse `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type OpenGradeCreateResponse struct {
	DeptId int64 `json:"dept_id,omitempty"`
}
