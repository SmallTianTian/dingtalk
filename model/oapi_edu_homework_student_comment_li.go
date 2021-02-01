package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkStudentCommentListRequest() *OapiEduHomeworkStudentCommentListRequest {
	return &OapiEduHomeworkStudentCommentListRequest{}
}

type OapiEduHomeworkStudentCommentListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkStudentCommentListResponse
	BizCode         string
	ClassId         int64
	HwId            int64
	StudentId       string
	TeacherUserid   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduHomeworkStudentCommentListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetStudentId(studentId2 string) {
	this.StudentId = studentId2
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetStudentId() string {
	return this.StudentId
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.comment.list"
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentCommentListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams["hw_id"] = this.HwId
	txtParams["student_id"] = this.StudentId
	txtParams["teacher_userid"] = this.TeacherUserid
	return txtParams
}
func (this *OapiEduHomeworkStudentCommentListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentCommentListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkStudentCommentListResponse struct {
	taobao.TaobaoResponse
	Result  []OpenHwCommentResponse `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type OpenHwCommentResponse struct {
	Attributes  string    `json:"attributes,omitempty"`
	Comment     string    `json:"comment,omitempty"`
	CommentId   int64     `json:"comment_id,omitempty"`
	CommentTime time.Time `json:"comment_time,omitempty"`
	Media       string    `json:"media,omitempty"`
	Photo       string    `json:"photo,omitempty"`
	StudentId   string    `json:"student_id,omitempty"`
	StudentName string    `json:"student_name,omitempty"`
	TeacherId   string    `json:"teacher_id,omitempty"`
	TeacherName string    `json:"teacher_name,omitempty"`
	Video       string    `json:"video,omitempty"`
}
