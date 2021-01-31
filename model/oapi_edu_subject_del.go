package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSubjectDeleteRequest() *OapiEduSubjectDeleteRequest {
	return &OapiEduSubjectDeleteRequest{}
}

type OapiEduSubjectDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSubjectDeleteResponse
	OperatorUserid  string
	PeriodCode      string
	SubjectCode     string
	SubjectName     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSubjectDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSubjectDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSubjectDeleteRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiEduSubjectDeleteRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiEduSubjectDeleteRequest) SetPeriodCode(periodCode2 string) {
	this.PeriodCode = periodCode2
}
func (this *OapiEduSubjectDeleteRequest) GetPeriodCode() string {
	return this.PeriodCode
}
func (this *OapiEduSubjectDeleteRequest) SetSubjectCode(subjectCode2 string) {
	this.SubjectCode = subjectCode2
}
func (this *OapiEduSubjectDeleteRequest) GetSubjectCode() string {
	return this.SubjectCode
}
func (this *OapiEduSubjectDeleteRequest) SetSubjectName(subjectName2 string) {
	this.SubjectName = subjectName2
}
func (this *OapiEduSubjectDeleteRequest) GetSubjectName() string {
	return this.SubjectName
}
func (this *OapiEduSubjectDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.subject.delete"
}
func (this *OapiEduSubjectDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSubjectDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSubjectDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSubjectDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSubjectDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["period_code"] = this.PeriodCode
	txtParams["subject_code"] = this.SubjectCode
	txtParams["subject_name"] = this.SubjectName
	return txtParams
}
func (this *OapiEduSubjectDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSubjectDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSubjectDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduSubjectDeleteResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
