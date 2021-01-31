package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduGradeQueryRequest() *OapiEduGradeQueryRequest {
	return &OapiEduGradeQueryRequest{}
}

type OapiEduGradeQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduGradeQueryResponse
	CampusId        int64
	Operator        string
	PeriodId        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduGradeQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduGradeQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduGradeQueryRequest) SetCampusId(campusId2 int64) {
	this.CampusId = campusId2
}
func (this *OapiEduGradeQueryRequest) GetCampusId() int64 {
	return this.CampusId
}
func (this *OapiEduGradeQueryRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduGradeQueryRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduGradeQueryRequest) SetPeriodId(periodId2 int64) {
	this.PeriodId = periodId2
}
func (this *OapiEduGradeQueryRequest) GetPeriodId() int64 {
	return this.PeriodId
}
func (this *OapiEduGradeQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.grade.query"
}
func (this *OapiEduGradeQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduGradeQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduGradeQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduGradeQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduGradeQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["campus_id"] = this.CampusId
	txtParams["operator"] = this.Operator
	txtParams["period_id"] = this.PeriodId
	return txtParams
}
func (this *OapiEduGradeQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Operator, "operator"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduGradeQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduGradeQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduGradeQueryResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type Grades struct {
	CampusId   int64  `json:"campus_id,omitempty"`
	DeptId     int64  `json:"dept_id,omitempty"`
	Grade      int64  `json:"grade,omitempty"`
	Name       string `json:"name,omitempty"`
	PeriodCode string `json:"period_code,omitempty"`
	SuperId    int64  `json:"super_id,omitempty"`
}
