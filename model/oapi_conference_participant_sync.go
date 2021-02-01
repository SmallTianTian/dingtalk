package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiConferenceParticipantSyncRequest() *OapiConferenceParticipantSyncRequest {
	return &OapiConferenceParticipantSyncRequest{}
}

type OapiConferenceParticipantSyncRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiConferenceParticipantSyncResponse
	BatchId               string
	BatchIndex            int64
	ConferenceId          string
	IsFinished            bool
	ParticipantUseridList string
	TopHttpMethod         string
	TopResponseType       string
	Userid                string
}

func (this *OapiConferenceParticipantSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConferenceParticipantSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConferenceParticipantSyncRequest) SetBatchId(batchId2 string) {
	this.BatchId = batchId2
}
func (this *OapiConferenceParticipantSyncRequest) GetBatchId() string {
	return this.BatchId
}
func (this *OapiConferenceParticipantSyncRequest) SetBatchIndex(batchIndex2 int64) {
	this.BatchIndex = batchIndex2
}
func (this *OapiConferenceParticipantSyncRequest) GetBatchIndex() int64 {
	return this.BatchIndex
}
func (this *OapiConferenceParticipantSyncRequest) SetConferenceId(conferenceId2 string) {
	this.ConferenceId = conferenceId2
}
func (this *OapiConferenceParticipantSyncRequest) GetConferenceId() string {
	return this.ConferenceId
}
func (this *OapiConferenceParticipantSyncRequest) SetIsFinished(isFinished2 bool) {
	this.IsFinished = isFinished2
}
func (this *OapiConferenceParticipantSyncRequest) GetIsFinished() bool {
	return this.IsFinished
}
func (this *OapiConferenceParticipantSyncRequest) SetParticipantUseridList(participantUseridList2 string) {
	this.ParticipantUseridList = participantUseridList2
}
func (this *OapiConferenceParticipantSyncRequest) GetParticipantUseridList() string {
	return this.ParticipantUseridList
}
func (this *OapiConferenceParticipantSyncRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiConferenceParticipantSyncRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiConferenceParticipantSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.conference.participant.sync"
}
func (this *OapiConferenceParticipantSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConferenceParticipantSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConferenceParticipantSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConferenceParticipantSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConferenceParticipantSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["batch_id"] = this.BatchId
	txtParams["batch_index"] = this.BatchIndex
	txtParams["conference_id"] = this.ConferenceId
	txtParams["is_finished"] = this.IsFinished
	txtParams["participant_userid_list"] = this.ParticipantUseridList
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiConferenceParticipantSyncRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BatchIndex, "batchIndex"); err != nil {
		return err
	}
	return nil
}
func (this *OapiConferenceParticipantSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConferenceParticipantSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiConferenceParticipantSyncResponse struct {
	taobao.TaobaoResponse
	BatchId string `json:"batch_id,omitempty"`
}
