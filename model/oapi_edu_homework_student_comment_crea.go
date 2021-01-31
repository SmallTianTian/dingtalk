package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkStudentCommentCreateRequest() *OapiEduHomeworkStudentCommentCreateRequest {
	return &OapiEduHomeworkStudentCommentCreateRequest{}
}

type OapiEduHomeworkStudentCommentCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkStudentCommentCreateResponse
	Attributes      string
	BizCode         string
	ClassId         string
	Comment         string
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

func (this *OapiEduHomeworkStudentCommentCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetAttributes(attributes2 string) {
	this.Attributes = attributes2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetAttributes() string {
	return this.Attributes
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetClassId(classId2 string) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetClassId() string {
	return this.ClassId
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetComment(comment2 string) {
	this.Comment = comment2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetComment() string {
	return this.Comment
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetMedia(media2 string) {
	this.Media = media2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetMedia() string {
	return this.Media
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetPhoto(photo2 string) {
	this.Photo = photo2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetPhoto() string {
	return this.Photo
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetStudentId(studentId2 string) {
	this.StudentId = studentId2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetStudentId() string {
	return this.StudentId
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetStudentName(studentName2 string) {
	this.StudentName = studentName2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetStudentName() string {
	return this.StudentName
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetVideo(video2 string) {
	this.Video = video2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetVideo() string {
	return this.Video
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.comment.create"
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attributes"] = this.Attributes
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams["comment"] = this.Comment
	txtParams["hw_id"] = this.HwId
	txtParams["media"] = this.Media
	txtParams["photo"] = this.Photo
	txtParams["student_id"] = this.StudentId
	txtParams["student_name"] = this.StudentName
	txtParams["teacher_userid"] = this.TeacherUserid
	txtParams["video"] = this.Video
	return txtParams
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentCommentCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkStudentCommentCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  int64  `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
