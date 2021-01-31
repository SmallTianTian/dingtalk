package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduStudentGetRequest() *OapiEduStudentGetRequest {
	return &OapiEduStudentGetRequest{}
}

type OapiEduStudentGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduStudentGetResponse
	ClassId         int64
	StudentUserid   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduStudentGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduStudentGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduStudentGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduStudentGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduStudentGetRequest) SetStudentUserid(studentUserid2 string) {
	this.StudentUserid = studentUserid2
}
func (this *OapiEduStudentGetRequest) GetStudentUserid() string {
	return this.StudentUserid
}
func (this *OapiEduStudentGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.student.get"
}
func (this *OapiEduStudentGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduStudentGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduStudentGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduStudentGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduStudentGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["student_userid"] = this.StudentUserid
	return txtParams
}
func (this *OapiEduStudentGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduStudentGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduStudentGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduStudentGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  StudentRespone `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type StudentRespone struct {
	CampusId  int64             `json:"campus_id,omitempty"`
	ClassId   int64             `json:"class_id,omitempty"`
	GradeId   int64             `json:"grade_id,omitempty"`
	Guardians []GuardianRespone `json:"guardians,omitempty"`
	Name      string            `json:"name,omitempty"`
	PeriodId  int64             `json:"period_id,omitempty"`
	StudentNo string            `json:"student_no,omitempty"`
}
