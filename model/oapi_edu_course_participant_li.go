package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduCourseParticipantListRequest() *OapiEduCourseParticipantListRequest {
	return &OapiEduCourseParticipantListRequest{}
}

type OapiEduCourseParticipantListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCourseParticipantListResponse
	CourseCode      string
	Cursor          int64
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCourseParticipantListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCourseParticipantListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCourseParticipantListRequest) SetCourseCode(courseCode2 string) {
	this.CourseCode = courseCode2
}
func (this *OapiEduCourseParticipantListRequest) GetCourseCode() string {
	return this.CourseCode
}
func (this *OapiEduCourseParticipantListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiEduCourseParticipantListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiEduCourseParticipantListRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiEduCourseParticipantListRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiEduCourseParticipantListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiEduCourseParticipantListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiEduCourseParticipantListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.course.participant.list"
}
func (this *OapiEduCourseParticipantListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCourseParticipantListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCourseParticipantListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCourseParticipantListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCourseParticipantListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["course_code"] = this.CourseCode
	txtParams["cursor"] = this.Cursor
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiEduCourseParticipantListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CourseCode, "courseCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduCourseParticipantListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCourseParticipantListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCourseParticipantListResponse struct {
	taobao.TaobaoResponse
	Result  ListCourseParticipantResponse `json:"result,omitempty"`
	Success bool                          `json:"success,omitempty"`
}
type CourseParticipantVO struct {
	ParticipantCorpid string `json:"participant_corpid,omitempty"`
	ParticipantId     string `json:"participant_id,omitempty"`
	ParticipantType   string `json:"participant_type,omitempty"`
	Role              string `json:"role,omitempty"`
}
type ListCourseParticipantResponse struct {
	HasMore    bool                  `json:"has_more,omitempty"`
	List       []CourseParticipantVO `json:"list,omitempty"`
	NextCursor int64                 `json:"next_cursor,omitempty"`
}
