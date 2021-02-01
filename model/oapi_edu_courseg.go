package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseGetRequest() *OapiEduCourseGetRequest {
	return &OapiEduCourseGetRequest{}
}

type OapiEduCourseGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseGetResponse
	CourseCode      string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseGetRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseGetRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseGetRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseGetRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.get"
}
func (this *OapiEduCourseGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiEduCourseGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseGetResponse struct {
	taobao.TaobaoResponse
	Result  Course `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type Course struct {
	BizKey        string `json:"biz_key,omitempty"`
	Code          string `json:"code,omitempty"`
	EndTime       int64  `json:"end_time,omitempty"`
	Introduce     string `json:"introduce,omitempty"`
	Name          string `json:"name,omitempty"`
	StartTime     int64  `json:"start_time,omitempty"`
	TeacherCorpid string `json:"teacher_corpid,omitempty"`
	TeacherUserid string `json:"teacher_userid,omitempty"`
}
