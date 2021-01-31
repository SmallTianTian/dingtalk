package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseReplayRequest() *OapiEduCourseReplayRequest {
	return &OapiEduCourseReplayRequest{}
}

type OapiEduCourseReplayRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseReplayResponse
	CourseCode      string
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseReplayRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseReplayRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseReplayRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseReplayRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseReplayRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiEduCourseReplayRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiEduCourseReplayRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.replay"
}
func (this *OapiEduCourseReplayRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseReplayRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseReplayRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseReplayRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseReplayRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiEduCourseReplayRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseReplayRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseReplayRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseReplayResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  ReplayCourseResponse `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type ReplayCourseResponse struct {
	ReplayUrls []string `json:"replay_urls,omitempty"`
	Replayable bool     `json:"replayable,omitempty"`
}
