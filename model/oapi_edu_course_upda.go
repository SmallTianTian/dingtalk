package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseUpdateRequest() *OapiEduCourseUpdateRequest {
	return &OapiEduCourseUpdateRequest{}
}

type OapiEduCourseUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseUpdateResponse
	CourseCode      string
	EndTime         int64
	Introduce       string
	Name            string
	OpUserid        string
	StartTime       int64
	TeacherCorpid   string
	TeacherUserid   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseUpdateRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseUpdateRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseUpdateRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiEduCourseUpdateRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiEduCourseUpdateRequest) SetIntroduce(introduce2 string) {
	this.Introduce = introduce2
}
func (this *OapiEduCourseUpdateRequest) GetIntroduce() string {
	return this.Introduce
}
func (this *OapiEduCourseUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiEduCourseUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiEduCourseUpdateRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseUpdateRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseUpdateRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiEduCourseUpdateRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiEduCourseUpdateRequest) SetTeacherCorpid(teacherCorpid2 string) {
	this.TeacherCorpid = teacherCorpid2
}
func (this *OapiEduCourseUpdateRequest) GetTeacherCorpid() string {
	return this.TeacherCorpid
}
func (this *OapiEduCourseUpdateRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduCourseUpdateRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduCourseUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.update"
}
func (this *OapiEduCourseUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["end_time"] = this.EndTime
	txtParams["introduce"] = this.Introduce
	txtParams["name"] = this.Name
	txtParams["op_userid"] = this.OpUserid
	txtParams["start_time"] = this.StartTime
	txtParams["teacher_corpid"] = this.TeacherCorpid
	txtParams["teacher_userid"] = this.TeacherUserid
	return txtParams
}
func (this *OapiEduCourseUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseUpdateResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
