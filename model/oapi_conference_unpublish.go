package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiConferenceUnpublishRequest() *OapiConferenceUnpublishRequest {
	return &OapiConferenceUnpublishRequest{}
}

type OapiConferenceUnpublishRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiConferenceUnpublishResponse
	ConferenceId    string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiConferenceUnpublishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConferenceUnpublishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConferenceUnpublishRequest) SetConferenceId(conferenceId2 string) {
	this.ConferenceId = conferenceId2
}
func (this *OapiConferenceUnpublishRequest) GetConferenceId() string {
	return this.ConferenceId
}
func (this *OapiConferenceUnpublishRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiConferenceUnpublishRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiConferenceUnpublishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.conference.unpublish"
}
func (this *OapiConferenceUnpublishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConferenceUnpublishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConferenceUnpublishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConferenceUnpublishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConferenceUnpublishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conference_id"] = this.ConferenceId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiConferenceUnpublishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConferenceId, "conferenceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiConferenceUnpublishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConferenceUnpublishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiConferenceUnpublishResponse struct {
	taobao.TaobaoResponse
}
