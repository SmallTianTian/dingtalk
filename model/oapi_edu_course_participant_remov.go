package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseParticipantRemoveRequest() *OapiEduCourseParticipantRemoveRequest {
	return &OapiEduCourseParticipantRemoveRequest{}
}

type OapiEduCourseParticipantRemoveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiEduCourseParticipantRemoveResponse
	CourseCode        string
	OpUserid          string
	ParticipantCorpid string
	ParticipantId     string
	ParticipantType   int64
	Role              string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiEduCourseParticipantRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseParticipantRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseParticipantRemoveRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseParticipantRemoveRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseParticipantRemoveRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseParticipantRemoveRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseParticipantRemoveRequest) SetParticipantCorpid(participantCorpid2 string) {
	this.ParticipantCorpid = participantCorpid2
}
func (this *OapiEduCourseParticipantRemoveRequest) GetParticipantCorpid() string {
	return this.ParticipantCorpid
}
func (this *OapiEduCourseParticipantRemoveRequest) SetParticipantId(participantId2 string) {
	this.ParticipantId = participantId2
}
func (this *OapiEduCourseParticipantRemoveRequest) GetParticipantId() string {
	return this.ParticipantId
}
func (this *OapiEduCourseParticipantRemoveRequest) SetParticipantType(participantType2 int64) {
	this.ParticipantType = participantType2
}
func (this *OapiEduCourseParticipantRemoveRequest) GetParticipantType() int64 {
	return this.ParticipantType
}
func (this *OapiEduCourseParticipantRemoveRequest) SetRole(role2 string) {
	this.Role = role2
}
func (this *OapiEduCourseParticipantRemoveRequest) GetRole() string {
	return this.Role
}
func (this *OapiEduCourseParticipantRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.participant.remove"
}
func (this *OapiEduCourseParticipantRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseParticipantRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseParticipantRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseParticipantRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseParticipantRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["op_userid"] = this.OpUserid
	txtParams["participant_corpid"] = this.ParticipantCorpid
	txtParams["participant_id"] = this.ParticipantId
	txtParams["participant_type"] = this.ParticipantType
	txtParams["role"] = this.Role
	return txtParams
}
func (this *OapiEduCourseParticipantRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseParticipantRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseParticipantRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseParticipantRemoveResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
