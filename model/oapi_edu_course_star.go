package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseStartRequest() *OapiEduCourseStartRequest {
	return &OapiEduCourseStartRequest{}
}

type OapiEduCourseStartRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseStartResponse
	CourseCode      string
	OpUserId        string
	StartOption     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseStartRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseStartRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseStartRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseStartRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseStartRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiEduCourseStartRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiEduCourseStartRequest) SetStartOption(startOption2 string) {
	this.StartOption = startOption2
}
func (this *OapiEduCourseStartRequest) GetStartOption() string {
	return this.StartOption
}
func (this *OapiEduCourseStartRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.start"
}
func (this *OapiEduCourseStartRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseStartRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseStartRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseStartRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseStartRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_user_id"] = this.OpUserId
	txtParams["start_option"] = this.StartOption
	return txtParams
}
func (this *OapiEduCourseStartRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseStartRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseStartRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type StartOption struct {
	BAllowJoinInAdvance bool `json:"b_allow_join_in_advance,omitempty"`
}
type OapiEduCourseStartResponse struct {
	taobao.TaobaoResponse
	Result  StartCourseResponse `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
type StartCourseResponse struct {
	IsReuse    bool   `json:"is_reuse,omitempty"`
	TargetId   string `json:"target_id,omitempty"`
	TargetType int64  `json:"target_type,omitempty"`
}
