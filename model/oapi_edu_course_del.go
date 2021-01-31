package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseDeleteRequest() *OapiEduCourseDeleteRequest {
	return &OapiEduCourseDeleteRequest{}
}

type OapiEduCourseDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseDeleteResponse
	CourseCode      string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseDeleteRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseDeleteRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseDeleteRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseDeleteRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.delete"
}
func (this *OapiEduCourseDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiEduCourseDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseDeleteResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
