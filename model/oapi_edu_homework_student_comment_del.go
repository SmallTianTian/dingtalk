package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduHomeworkStudentCommentDeleteRequest() *OapiEduHomeworkStudentCommentDeleteRequest {
	return &OapiEduHomeworkStudentCommentDeleteRequest{}
}

type OapiEduHomeworkStudentCommentDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkStudentCommentDeleteResponse
	BizCode         string
	ClassId         int64
	CommentId       int64
	HwId            int64
	StudentId       string
	TeacherUserid   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetCommentId(commentId2 int64) {
	this.CommentId = commentId2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetCommentId() int64 {
	return this.CommentId
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetStudentId(studentId2 string) {
	this.StudentId = studentId2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetStudentId() string {
	return this.StudentId
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.comment.delete"
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams["comment_id"] = this.CommentId
	txtParams["hw_id"] = this.HwId
	txtParams["student_id"] = this.StudentId
	txtParams["teacher_userid"] = this.TeacherUserid
	return txtParams
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentCommentDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkStudentCommentDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
