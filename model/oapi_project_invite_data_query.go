package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProjectInviteDataQueryRequest() *OapiProjectInviteDataQueryRequest {
	return &OapiProjectInviteDataQueryRequest{}
}

type OapiProjectInviteDataQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProjectInviteDataQueryResponse
	InviteDataQuery string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProjectInviteDataQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProjectInviteDataQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProjectInviteDataQueryRequest) SetInviteDataQuery(inviteDataQuery2 string) {
	this.InviteDataQuery = inviteDataQuery2
}
func (this *OapiProjectInviteDataQueryRequest) GetInviteDataQuery() string {
	return this.InviteDataQuery
}
func (this *OapiProjectInviteDataQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.project.invite.data.query"
}
func (this *OapiProjectInviteDataQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProjectInviteDataQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProjectInviteDataQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProjectInviteDataQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProjectInviteDataQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["invite_data_query"] = this.InviteDataQuery
	return txtParams
}
func (this *OapiProjectInviteDataQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProjectInviteDataQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProjectInviteDataQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type InviteDataQuery struct {
	Cursor int64 `json:"cursor,omitempty"`
	Size   int64 `json:"size,omitempty"`
	Status int64 `json:"status,omitempty"`
}
type OapiProjectInviteDataQueryResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type InviteDataModel struct {
	Channel      string `json:"channel,omitempty"`
	CorpId       string `json:"corp_id,omitempty"`
	Extension    string `json:"extension,omitempty"`
	GmtModified  int64  `json:"gmt_modified,omitempty"`
	InviteUserid string `json:"invite_userid,omitempty"`
	JoinAt       int64  `json:"join_at,omitempty"`
	Status       int64  `json:"status,omitempty"`
	Userid       string `json:"userid,omitempty"`
}
