package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiConferenceParticipantDeleteRequest() *OapiConferenceParticipantDeleteRequest {
	return &OapiConferenceParticipantDeleteRequest{}
}

type OapiConferenceParticipantDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiConferenceParticipantDeleteResponse
	ConferenceId          string
	ParticipantUseridList string
	TopHttpMethod         string
	TopResponseType       string
	Userid                string
}

func (this *OapiConferenceParticipantDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConferenceParticipantDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConferenceParticipantDeleteRequest) SetConferenceId(conferenceId2 string) {
	this.ConferenceId = conferenceId2
}
func (this *OapiConferenceParticipantDeleteRequest) GetConferenceId() string {
	return this.ConferenceId
}
func (this *OapiConferenceParticipantDeleteRequest) SetParticipantUseridList(participantUseridList2 string) {
	this.ParticipantUseridList = participantUseridList2
}
func (this *OapiConferenceParticipantDeleteRequest) GetParticipantUseridList() string {
	return this.ParticipantUseridList
}
func (this *OapiConferenceParticipantDeleteRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiConferenceParticipantDeleteRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiConferenceParticipantDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.conference.participant.delete"
}
func (this *OapiConferenceParticipantDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConferenceParticipantDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConferenceParticipantDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConferenceParticipantDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConferenceParticipantDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conference_id"] = this.ConferenceId
	txtParams["participant_userid_list"] = this.ParticipantUseridList
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiConferenceParticipantDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConferenceId, "conferenceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiConferenceParticipantDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConferenceParticipantDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiConferenceParticipantDeleteResponse struct {
	taobao.TaobaoResponse
}
