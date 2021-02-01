package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSubjectCreateRequest() *OapiEduSubjectCreateRequest {
	return &OapiEduSubjectCreateRequest{}
}

type OapiEduSubjectCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSubjectCreateResponse
	OperatorUserid  string
	PeriodCode      string
	SubjectCode     string
	SubjectName     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSubjectCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSubjectCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSubjectCreateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiEduSubjectCreateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiEduSubjectCreateRequest) SetPeriodCode(periodCode2 string) {
	this.PeriodCode = periodCode2
}
func (this *OapiEduSubjectCreateRequest) GetPeriodCode() string {
	return this.PeriodCode
}
func (this *OapiEduSubjectCreateRequest) SetSubjectCode(subjectCode2 string) {
	this.SubjectCode = subjectCode2
}
func (this *OapiEduSubjectCreateRequest) GetSubjectCode() string {
	return this.SubjectCode
}
func (this *OapiEduSubjectCreateRequest) SetSubjectName(subjectName2 string) {
	this.SubjectName = subjectName2
}
func (this *OapiEduSubjectCreateRequest) GetSubjectName() string {
	return this.SubjectName
}
func (this *OapiEduSubjectCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.subject.create"
}
func (this *OapiEduSubjectCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSubjectCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSubjectCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSubjectCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSubjectCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["period_code"] = this.PeriodCode
	txtParams["subject_code"] = this.SubjectCode
	txtParams["subject_name"] = this.SubjectName
	return txtParams
}
func (this *OapiEduSubjectCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSubjectCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSubjectCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduSubjectCreateResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
