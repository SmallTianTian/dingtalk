package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseCreateRequest() *OapiEduCourseCreateRequest {
	return &OapiEduCourseCreateRequest{}
}

type OapiEduCourseCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseCreateResponse
	BizKey          string
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

func (this *OapiEduCourseCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseCreateRequest) SetBizKey(bizKey2 string) {
	this.BizKey = bizKey2
}
func (this *OapiEduCourseCreateRequest) GetBizKey() string {
	return this.BizKey
}
func (this *OapiEduCourseCreateRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiEduCourseCreateRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiEduCourseCreateRequest) SetIntroduce(introduce2 string) {
	this.Introduce = introduce2
}
func (this *OapiEduCourseCreateRequest) GetIntroduce() string {
	return this.Introduce
}
func (this *OapiEduCourseCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiEduCourseCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiEduCourseCreateRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseCreateRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseCreateRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiEduCourseCreateRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiEduCourseCreateRequest) SetTeacherCorpid(teacherCorpid2 string) {
	this.TeacherCorpid = teacherCorpid2
}
func (this *OapiEduCourseCreateRequest) GetTeacherCorpid() string {
	return this.TeacherCorpid
}
func (this *OapiEduCourseCreateRequest) SetTeacherUserid(teacherUserid2 string) {
	this.TeacherUserid = teacherUserid2
}
func (this *OapiEduCourseCreateRequest) GetTeacherUserid() string {
	return this.TeacherUserid
}
func (this *OapiEduCourseCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.create"
}
func (this *OapiEduCourseCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_key"] = this.BizKey
	txtParams["end_time"] = this.EndTime
	txtParams["introduce"] = this.Introduce
	txtParams["name"] = this.Name
	txtParams["op_userid"] = this.OpUserid
	txtParams["start_time"] = this.StartTime
	txtParams["teacher_corpid"] = this.TeacherCorpid
	txtParams["teacher_userid"] = this.TeacherUserid
	return txtParams
}
func (this *OapiEduCourseCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizKey, "bizKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseCreateResponse struct {
	taobao.TaobaoResponse
	CourseCode string `json:"course_code,omitempty"`

	Success bool `json:"success,omitempty"`
}
