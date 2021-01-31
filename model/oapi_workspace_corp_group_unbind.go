package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceCorpGroupUnbindRequest() *OapiWorkspaceCorpGroupUnbindRequest {
	return &OapiWorkspaceCorpGroupUnbindRequest{}
}

type OapiWorkspaceCorpGroupUnbindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCorpGroupUnbindResponse
	ChatIds         string
	TargetCorpId    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceCorpGroupUnbindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) SetChatIds(chatIds2 string) {
	this.ChatIds = chatIds2
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetChatIds() string {
	return this.ChatIds
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) SetTargetCorpId(targetCorpId2 string) {
	this.TargetCorpId = targetCorpId2
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetTargetCorpId() string {
	return this.TargetCorpId
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.corp.group.unbind"
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_ids"] = this.ChatIds
	txtParams["target_corp_id"] = this.TargetCorpId
	return txtParams
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatIds, "chatIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCorpGroupUnbindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceCorpGroupUnbindResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Result  []string `json:"result,omitempty"`
}
