package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSceneservicegroupGroupGetRequest() *OapiSceneservicegroupGroupGetRequest {
	return &OapiSceneservicegroupGroupGetRequest{}
}

type OapiSceneservicegroupGroupGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiSceneservicegroupGroupGetResponse
	OpenConversationid string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiSceneservicegroupGroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSceneservicegroupGroupGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSceneservicegroupGroupGetRequest) SetOpenConversationid(openConversationid2 string) {
	this.OpenConversationid = openConversationid2
}
func (this *OapiSceneservicegroupGroupGetRequest) GetOpenConversationid() string {
	return this.OpenConversationid
}
func (this *OapiSceneservicegroupGroupGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sceneservicegroup.group.get"
}
func (this *OapiSceneservicegroupGroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSceneservicegroupGroupGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSceneservicegroupGroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSceneservicegroupGroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSceneservicegroupGroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_conversationid"] = this.OpenConversationid
	return txtParams
}
func (this *OapiSceneservicegroupGroupGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationid, "openConversationid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSceneservicegroupGroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSceneservicegroupGroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSceneservicegroupGroupGetResponse struct {
	taobao.TaobaoResponse
	Result GetGroupResponse `json:"result,omitempty"`
}
type GetGroupResponse struct {
	GroupName          string `json:"group_name,omitempty"`
	GroupUrl           string `json:"group_url,omitempty"`
	OpenConversationid string `json:"open_conversationid,omitempty"`
	OpenGroupsetid     string `json:"open_groupsetid,omitempty"`
	OpenTeamid         string `json:"open_teamid,omitempty"`
}
