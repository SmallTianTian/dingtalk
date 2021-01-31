package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkStudentCommentUpdateRequest() *OapiEduHomeworkStudentCommentUpdateRequest {
	return &OapiEduHomeworkStudentCommentUpdateRequest{}
}

type OapiEduHomeworkStudentCommentUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkStudentCommentUpdateResponse
	Attributes      string
	BizCode         string
	ClassId         int64
	Comment         string
	CommentId       int64
	HwId            int64
	Media           string
	Photo           string
	StudentId       string
	StudentName     string
	TeacherUserid   string
	TopHttpMethod   string
	TopResponseType string
	Video           string
}

func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetAttributes(attributes2 string) {
	this.Attributes = attributes2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetAttributes() string {
	return this.Attributes
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetComment(comment2 string) {
	this.Comment = comment2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetComment() string {
	return this.Comment
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetCommentId(commentId2 int64) {
	this.CommentId = commentId2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetCommentId() int64 {
	return this.CommentId
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetMedia(media2 string) {
	this.Media = media2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetMedia() string {
	return this.Media
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetPhoto(photo2 string) {
	this.Photo = photo2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetPhoto() string {
	return this.Photo
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetStudentId(studentId2 string) {
	this.StudentId = studentId2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetStudentId() string {
	return this.StudentId
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetStudentName(studentName2 string) {
	this.StudentName = studentName2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetStudentName() string {
	return this.StudentName
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetVideo(video2 string) {
	this.Video = video2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetVideo() string {
	return this.Video
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.comment.update"
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attributes"] = this.Attributes
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams["comment"] = this.Comment
	txtParams["comment_id"] = this.CommentId
	txtParams["hw_id"] = this.HwId
	txtParams["media"] = this.Media
	txtParams["photo"] = this.Photo
	txtParams["student_id"] = this.StudentId
	txtParams["student_name"] = this.StudentName
	txtParams["teacher_userid"] = this.TeacherUserid
	txtParams["video"] = this.Video
	return txtParams
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentCommentUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkStudentCommentUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  int64  `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
