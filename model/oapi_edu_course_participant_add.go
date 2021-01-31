package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseParticipantAddRequest() *OapiEduCourseParticipantAddRequest {
	return &OapiEduCourseParticipantAddRequest{}
}

type OapiEduCourseParticipantAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiEduCourseParticipantAddResponse
	CourseCode        string
	OpUserid          string
	ParticipantCorpid string
	ParticipantId     string
	ParticipantType   int64
	Role              string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiEduCourseParticipantAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseParticipantAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseParticipantAddRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseParticipantAddRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseParticipantAddRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseParticipantAddRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseParticipantAddRequest) SetParticipantCorpid(participantCorpid2 string) {
	this.ParticipantCorpid = participantCorpid2
}
func (this *OapiEduCourseParticipantAddRequest) GetParticipantCorpid() string {
	return this.ParticipantCorpid
}
func (this *OapiEduCourseParticipantAddRequest) SetParticipantId(participantId2 string) {
	this.ParticipantId = participantId2
}
func (this *OapiEduCourseParticipantAddRequest) GetParticipantId() string {
	return this.ParticipantId
}
func (this *OapiEduCourseParticipantAddRequest) SetParticipantType(participantType2 int64) {
	this.ParticipantType = participantType2
}
func (this *OapiEduCourseParticipantAddRequest) GetParticipantType() int64 {
	return this.ParticipantType
}
func (this *OapiEduCourseParticipantAddRequest) SetRole(role2 string) {
	this.Role = role2
}
func (this *OapiEduCourseParticipantAddRequest) GetRole() string {
	return this.Role
}
func (this *OapiEduCourseParticipantAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.participant.add"
}
func (this *OapiEduCourseParticipantAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseParticipantAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseParticipantAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseParticipantAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseParticipantAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_userid"] = this.OpUserid
	txtParams["participant_corpid"] = this.ParticipantCorpid
	txtParams["participant_id"] = this.ParticipantId
	txtParams["participant_type"] = this.ParticipantType
	txtParams["role"] = this.Role
	return txtParams
}
func (this *OapiEduCourseParticipantAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseParticipantAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseParticipantAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseParticipantAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
