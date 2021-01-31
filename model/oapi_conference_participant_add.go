package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiConferenceParticipantAddRequest() *OapiConferenceParticipantAddRequest {
	return &OapiConferenceParticipantAddRequest{}
}

type OapiConferenceParticipantAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiConferenceParticipantAddResponse
	ConferenceId          string
	ParticipantUseridList string
	TopHttpMethod         string
	TopResponseType       string
	Userid                string
}

func (this *OapiConferenceParticipantAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConferenceParticipantAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConferenceParticipantAddRequest) SetConferenceId(conferenceId2 string) {
	this.ConferenceId = conferenceId2
}
func (this *OapiConferenceParticipantAddRequest) GetConferenceId() string {
	return this.ConferenceId
}
func (this *OapiConferenceParticipantAddRequest) SetParticipantUseridList(participantUseridList2 string) {
	this.ParticipantUseridList = participantUseridList2
}
func (this *OapiConferenceParticipantAddRequest) GetParticipantUseridList() string {
	return this.ParticipantUseridList
}
func (this *OapiConferenceParticipantAddRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiConferenceParticipantAddRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiConferenceParticipantAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.conference.participant.add"
}
func (this *OapiConferenceParticipantAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConferenceParticipantAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConferenceParticipantAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConferenceParticipantAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConferenceParticipantAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conference_id"] = this.ConferenceId
	txtParams["participant_userid_list"] = this.ParticipantUseridList
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiConferenceParticipantAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConferenceId, "conferenceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiConferenceParticipantAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConferenceParticipantAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiConferenceParticipantAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
