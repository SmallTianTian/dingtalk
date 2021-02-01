package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiConferenceGetRequest() *OapiConferenceGetRequest {
	return &OapiConferenceGetRequest{}
}

type OapiConferenceGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiConferenceGetResponse
	ConferenceId    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiConferenceGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConferenceGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConferenceGetRequest) SetConferenceId(conferenceId2 string) {
	this.ConferenceId = conferenceId2
}
func (this *OapiConferenceGetRequest) GetConferenceId() string {
	return this.ConferenceId
}
func (this *OapiConferenceGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.conference.get"
}
func (this *OapiConferenceGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConferenceGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConferenceGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConferenceGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConferenceGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conference_id"] = this.ConferenceId
	return txtParams
}
func (this *OapiConferenceGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConferenceId, "conferenceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiConferenceGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConferenceGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiConferenceGetResponse struct {
	taobao.TaobaoResponse
	Result ConferenceInfoDo `json:"result,omitempty"`
}
type ConferenceInfoDo struct {
	Address           string   `json:"address,omitempty"`
	AdminUserid       string   `json:"admin_userid,omitempty"`
	ArrangeUseridList []string `json:"arrange_userid_list,omitempty"`
	ConferenceId      string   `json:"conference_id,omitempty"`
	Content           string   `json:"content,omitempty"`
	CorpId            string   `json:"corp_id,omitempty"`
	CreateUserid      string   `json:"create_userid,omitempty"`
	EndTime           int64    `json:"end_time,omitempty"`
	ModifiedUserid    string   `json:"modified_userid,omitempty"`
	Poi               string   `json:"poi,omitempty"`
	StartTime         int64    `json:"start_time,omitempty"`
	Status            int64    `json:"status,omitempty"`
	Topic             string   `json:"topic,omitempty"`
	TopicPicUrl       string   `json:"topic_pic_url,omitempty"`
	Type              int64    `json:"type,omitempty"`
	Version           int64    `json:"version,omitempty"`
}
