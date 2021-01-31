package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduHomeworkStudentMarkTagRequest() *OapiEduHomeworkStudentMarkTagRequest {
	return &OapiEduHomeworkStudentMarkTagRequest{}
}

type OapiEduHomeworkStudentMarkTagRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkStudentMarkTagResponse
	BizCode         string
	ClassId         int64
	HwId            int64
	StudentId       string
	StudentName     string
	Tag             string
	TeacherId       string
	Text            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduHomeworkStudentMarkTagRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetStudentId(studentId2 string) {
	this.StudentId = studentId2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetStudentId() string {
	return this.StudentId
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetStudentName(studentName2 string) {
	this.StudentName = studentName2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetStudentName() string {
	return this.StudentName
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetTag(tag2 string) {
	this.Tag = tag2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetTag() string {
	return this.Tag
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetTeacherId(teacherId2 string) {
	this.TeacherId = teacherId2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetTeacherId() string {
	return this.TeacherId
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetText(text2 string) {
	this.Text = text2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetText() string {
	return this.Text
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.mark.tag"
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentMarkTagRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams["hw_id"] = this.HwId
	txtParams["student_id"] = this.StudentId
	txtParams["student_name"] = this.StudentName
	txtParams["tag"] = this.Tag
	txtParams["teacher_id"] = this.TeacherId
	txtParams["text"] = this.Text
	return txtParams
}
func (this *OapiEduHomeworkStudentMarkTagRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentMarkTagRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkStudentMarkTagResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  int64  `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
