package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseEndRequest() *OapiEduCourseEndRequest {
	return &OapiEduCourseEndRequest{}
}

type OapiEduCourseEndRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseEndResponse
	CourseCode      string
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseEndRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseEndRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseEndRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseEndRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseEndRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiEduCourseEndRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiEduCourseEndRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.end"
}
func (this *OapiEduCourseEndRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseEndRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseEndRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseEndRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseEndRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiEduCourseEndRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseEndRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseEndRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseEndResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
