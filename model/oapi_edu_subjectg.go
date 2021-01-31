package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSubjectGetRequest() *OapiEduSubjectGetRequest {
	return &OapiEduSubjectGetRequest{}
}

type OapiEduSubjectGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSubjectGetResponse
	OperatorUserid  string
	PeriodCode      string
	SubjectCode     string
	SubjectName     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSubjectGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSubjectGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSubjectGetRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiEduSubjectGetRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiEduSubjectGetRequest) SetPeriodCode(periodCode2 string) {
	this.PeriodCode = periodCode2
}
func (this *OapiEduSubjectGetRequest) GetPeriodCode() string {
	return this.PeriodCode
}
func (this *OapiEduSubjectGetRequest) SetSubjectCode(subjectCode2 string) {
	this.SubjectCode = subjectCode2
}
func (this *OapiEduSubjectGetRequest) GetSubjectCode() string {
	return this.SubjectCode
}
func (this *OapiEduSubjectGetRequest) SetSubjectName(subjectName2 string) {
	this.SubjectName = subjectName2
}
func (this *OapiEduSubjectGetRequest) GetSubjectName() string {
	return this.SubjectName
}
func (this *OapiEduSubjectGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.subject.get"
}
func (this *OapiEduSubjectGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSubjectGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSubjectGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSubjectGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSubjectGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["period_code"] = this.PeriodCode
	txtParams["subject_code"] = this.SubjectCode
	txtParams["subject_name"] = this.SubjectName
	return txtParams
}
func (this *OapiEduSubjectGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSubjectGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSubjectGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduSubjectGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64              `json:"errcode,omitempty"`
	Errmsg  string             `json:"errmsg,omitempty"`
	Result  SubjectInstanceDTO `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type SubjectInstanceDTO struct {
	PeriodCode  string `json:"period_code,omitempty"`
	SubjectCode string `json:"subject_code,omitempty"`
	SubjectName string `json:"subject_name,omitempty"`
}
