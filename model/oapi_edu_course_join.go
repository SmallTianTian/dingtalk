package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseJoinRequest() *OapiEduCourseJoinRequest {
	return &OapiEduCourseJoinRequest{}
}

type OapiEduCourseJoinRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseJoinResponse
	CourseCode      string
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseJoinRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseJoinRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseJoinRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseJoinRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseJoinRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiEduCourseJoinRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiEduCourseJoinRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.join"
}
func (this *OapiEduCourseJoinRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseJoinRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseJoinRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseJoinRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseJoinRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiEduCourseJoinRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseJoinRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseJoinRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseJoinResponse struct {
	taobao.TaobaoResponse
	Result  JoinCourseResponse `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type JoinCourseResponse struct {
	JoinUrl  string `json:"join_url,omitempty"`
	Joinable bool   `json:"joinable,omitempty"`
}
