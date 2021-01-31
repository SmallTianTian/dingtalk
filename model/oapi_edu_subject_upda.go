package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduSubjectUpdateRequest() *OapiEduSubjectUpdateRequest {
	return &OapiEduSubjectUpdateRequest{}
}

type OapiEduSubjectUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduSubjectUpdateResponse
	OperatorUserid  string
	PeriodCode      string
	SubjectCode     string
	SubjectName     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduSubjectUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduSubjectUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduSubjectUpdateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiEduSubjectUpdateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiEduSubjectUpdateRequest) SetPeriodCode(periodCode2 string) {
	this.PeriodCode = periodCode2
}
func (this *OapiEduSubjectUpdateRequest) GetPeriodCode() string {
	return this.PeriodCode
}
func (this *OapiEduSubjectUpdateRequest) SetSubjectCode(subjectCode2 string) {
	this.SubjectCode = subjectCode2
}
func (this *OapiEduSubjectUpdateRequest) GetSubjectCode() string {
	return this.SubjectCode
}
func (this *OapiEduSubjectUpdateRequest) SetSubjectName(subjectName2 string) {
	this.SubjectName = subjectName2
}
func (this *OapiEduSubjectUpdateRequest) GetSubjectName() string {
	return this.SubjectName
}
func (this *OapiEduSubjectUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.subject.update"
}
func (this *OapiEduSubjectUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduSubjectUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduSubjectUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduSubjectUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduSubjectUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["period_code"] = this.PeriodCode
	txtParams["subject_code"] = this.SubjectCode
	txtParams["subject_name"] = this.SubjectName
	return txtParams
}
func (this *OapiEduSubjectUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduSubjectUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduSubjectUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduSubjectUpdateResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
