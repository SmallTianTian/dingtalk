package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiConferencePublishRequest() *OapiConferencePublishRequest {
	return &OapiConferencePublishRequest{}
}

type OapiConferencePublishRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiConferencePublishResponse
	ConferenceId    string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiConferencePublishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConferencePublishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConferencePublishRequest) SetConferenceId(conferenceId2 string) {
	this.ConferenceId = conferenceId2
}
func (this *OapiConferencePublishRequest) GetConferenceId() string {
	return this.ConferenceId
}
func (this *OapiConferencePublishRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiConferencePublishRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiConferencePublishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.conference.publish"
}
func (this *OapiConferencePublishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConferencePublishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConferencePublishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConferencePublishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConferencePublishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conference_id"] = this.ConferenceId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiConferencePublishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConferenceId, "conferenceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiConferencePublishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConferencePublishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiConferencePublishResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
