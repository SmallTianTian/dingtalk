package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkCreateRequest() *OapiEduHomeworkCreateRequest {
	return &OapiEduHomeworkCreateRequest{}
}

type OapiEduHomeworkCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiEduHomeworkCreateResponse
	Attributes       string
	BizCode          string
	CourseName       string
	HwContent        string
	HwDeadline       int64
	HwDeadlineOpen   string
	HwMedia          string
	HwPhoto          string
	HwTitle          string
	HwType           string
	HwVideo          string
	Identifier       string
	ScheduledRelease string
	ScheduledTime    string
	SelectClass      string
	SelectStu        string
	Status           string
	TargetRole       string
	TeacherName      string
	TeacherUserid    string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiEduHomeworkCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkCreateRequest) SetAttributes(attributes2 string) {
	this.Attributes = attributes2
}
func (this *OapiEduHomeworkCreateRequest) GetAttributes() string {
	return this.Attributes
}
func (this *OapiEduHomeworkCreateRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkCreateRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkCreateRequest) SetCourseName(courseName2 string) {
	this.CourseName = courseName2
}
func (this *OapiEduHomeworkCreateRequest) GetCourseName() string {
	return this.CourseName
}
func (this *OapiEduHomeworkCreateRequest) SetHwContent(hwContent2 string) {
	this.HwContent = hwContent2
}
func (this *OapiEduHomeworkCreateRequest) GetHwContent() string {
	return this.HwContent
}
func (this *OapiEduHomeworkCreateRequest) SetHwDeadline(hwDeadline2 int64) {
	this.HwDeadline = hwDeadline2
}
func (this *OapiEduHomeworkCreateRequest) GetHwDeadline() int64 {
	return this.HwDeadline
}
func (this *OapiEduHomeworkCreateRequest) SetHwDeadlineOpen(hwDeadlineOpen2 string) {
	this.HwDeadlineOpen = hwDeadlineOpen2
}
func (this *OapiEduHomeworkCreateRequest) GetHwDeadlineOpen() string {
	return this.HwDeadlineOpen
}
func (this *OapiEduHomeworkCreateRequest) SetHwMedia(hwMedia2 string) {
	this.HwMedia = hwMedia2
}
func (this *OapiEduHomeworkCreateRequest) GetHwMedia() string {
	return this.HwMedia
}
func (this *OapiEduHomeworkCreateRequest) SetHwPhoto(hwPhoto2 string) {
	this.HwPhoto = hwPhoto2
}
func (this *OapiEduHomeworkCreateRequest) GetHwPhoto() string {
	return this.HwPhoto
}
func (this *OapiEduHomeworkCreateRequest) SetHwTitle(hwTitle2 string) {
	this.HwTitle = hwTitle2
}
func (this *OapiEduHomeworkCreateRequest) GetHwTitle() string {
	return this.HwTitle
}
func (this *OapiEduHomeworkCreateRequest) SetHwType(hwType2 string) {
	this.HwType = hwType2
}
func (this *OapiEduHomeworkCreateRequest) GetHwType() string {
	return this.HwType
}
func (this *OapiEduHomeworkCreateRequest) SetHwVideo(hwVideo2 string) {
	this.HwVideo = hwVideo2
}
func (this *OapiEduHomeworkCreateRequest) GetHwVideo() string {
	return this.HwVideo
}
func (this *OapiEduHomeworkCreateRequest) SetIdentifier(identifier2 string) {
	this.Identifier = identifier2
}
func (this *OapiEduHomeworkCreateRequest) GetIdentifier() string {
	return this.Identifier
}
func (this *OapiEduHomeworkCreateRequest) SetScheduledRelease(scheduledRelease2 string) {
	this.ScheduledRelease = scheduledRelease2
}
func (this *OapiEduHomeworkCreateRequest) GetScheduledRelease() string {
	return this.ScheduledRelease
}
func (this *OapiEduHomeworkCreateRequest) SetScheduledTime(scheduledTime2 string) {
	this.ScheduledTime = scheduledTime2
}
func (this *OapiEduHomeworkCreateRequest) GetScheduledTime() string {
	return this.ScheduledTime
}
func (this *OapiEduHomeworkCreateRequest) SetSelectClass(selectClass2 string) {
	this.SelectClass = selectClass2
}
func (this *OapiEduHomeworkCreateRequest) GetSelectClass() string {
	return this.SelectClass
}
func (this *OapiEduHomeworkCreateRequest) SetSelectStu(selectStu2 string) {
	this.SelectStu = selectStu2
}
func (this *OapiEduHomeworkCreateRequest) GetSelectStu() string {
	return this.SelectStu
}
func (this *OapiEduHomeworkCreateRequest) SetStatus(status2 string) {
	this.Status = status2
}
func (this *OapiEduHomeworkCreateRequest) GetStatus() string {
	return this.Status
}
func (this *OapiEduHomeworkCreateRequest) SetTargetRole(targetRole2 string) {
	this.TargetRole = targetRole2
}
func (this *OapiEduHomeworkCreateRequest) GetTargetRole() string {
	return this.TargetRole
}
func (this *OapiEduHomeworkCreateRequest) SetTeacherName(teacherName2 string) {
	this.TeacherName = teacherName2
}
func (this *OapiEduHomeworkCreateRequest) GetTeacherName() string {
	return this.TeacherName
}
func (this *OapiEduHomeworkCreateRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduHomeworkCreateRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduHomeworkCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.create"
}
func (this *OapiEduHomeworkCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attributes"] = this.Attributes
	txtParams["biz_code"] = this.BizCode
	txtParams["course_name"] = this.CourseName
	txtParams["hw_content"] = this.HwContent
	txtParams["hw_deadline"] = this.HwDeadline
	txtParams["hw_deadline_open"] = this.HwDeadlineOpen
	txtParams["hw_media"] = this.HwMedia
	txtParams["hw_photo"] = this.HwPhoto
	txtParams["hw_title"] = this.HwTitle
	txtParams["hw_type"] = this.HwType
	txtParams["hw_video"] = this.HwVideo
	txtParams["identifier"] = this.Identifier
	txtParams["scheduled_release"] = this.ScheduledRelease
	txtParams["scheduled_time"] = this.ScheduledTime
	txtParams["select_class"] = this.SelectClass
	txtParams["select_stu"] = this.SelectStu
	txtParams["status"] = this.Status
	txtParams["target_role"] = this.TargetRole
	txtParams["teacher_name"] = this.TeacherName
	txtParams["teacher_userid"] = this.TeacherUserid
	return txtParams
}
func (this *OapiEduHomeworkCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ClassInfoItem struct {
	ClassId   int64  `json:"class_id,omitempty"`
	ClassName string `json:"class_name,omitempty"`
}
type StuInfoItem struct {
	StuId   string `json:"stu_id,omitempty"`
	StuName string `json:"stu_name,omitempty"`
}
type SelectStuItem struct {
	ClassId   int64         `json:"class_id,omitempty"`
	ClassName string        `json:"class_name,omitempty"`
	ClassStu  []StuInfoItem `json:"class_stu,omitempty"`
}
type OapiEduHomeworkCreateResponse struct {
	taobao.TaobaoResponse
	Result  int64 `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
