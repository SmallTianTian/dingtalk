package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkStudentSubmitRequest() *OapiEduHomeworkStudentSubmitRequest {
	return &OapiEduHomeworkStudentSubmitRequest{}
}

type OapiEduHomeworkStudentSubmitRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkStudentSubmitResponse
	Attributes      string
	BizCode         string
	ClassId         int64
	Content         string
	HwId            int64
	Media           string
	Photo           string
	StudentId       string
	StudentName     string
	SubmitorId      string
	Title           string
	TopHttpMethod   string
	TopResponseType string
	Video           string
}

func (this *OapiEduHomeworkStudentSubmitRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetAttributes(attributes2 string) {
	this.Attributes = attributes2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetAttributes() string {
	return this.Attributes
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetContent() string {
	return this.Content
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetHwId(hwId2 int64) {
	this.HwId = hwId2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetHwId() int64 {
	return this.HwId
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetMedia(media2 string) {
	this.Media = media2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetMedia() string {
	return this.Media
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetPhoto(photo2 string) {
	this.Photo = photo2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetPhoto() string {
	return this.Photo
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetStudentId(studentId2 string) {
	this.StudentId = studentId2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetStudentId() string {
	return this.StudentId
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetStudentName(studentName2 string) {
	this.StudentName = studentName2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetStudentName() string {
	return this.StudentName
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetSubmitorId(submitorId2 string) {
	this.SubmitorId = submitorId2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetSubmitorId() string {
	return this.SubmitorId
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetTitle() string {
	return this.Title
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetVideo(video2 string) {
	this.Video = video2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetVideo() string {
	return this.Video
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.student.submit"
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkStudentSubmitRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attributes"] = this.Attributes
	txtParams["biz_code"] = this.BizCode
	txtParams["class_id"] = this.ClassId
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["hw_id"] = this.HwId
	txtParams["media"] = this.Media
	txtParams["photo"] = this.Photo
	txtParams["student_id"] = this.StudentId
	txtParams["student_name"] = this.StudentName
	txtParams["submitor_id"] = this.SubmitorId
	txtParams["title"] = this.Title
	txtParams["video"] = this.Video
	return txtParams
}
func (this *OapiEduHomeworkStudentSubmitRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkStudentSubmitRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkStudentSubmitResponse struct {
	taobao.TaobaoResponse
	Result  int64 `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
