package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceCorpGroupBindRequest() *OapiWorkspaceCorpGroupBindRequest {
	return &OapiWorkspaceCorpGroupBindRequest{}
}

type OapiWorkspaceCorpGroupBindRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCorpGroupBindResponse
	ChatIds         string
	TargetCorpId    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceCorpGroupBindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCorpGroupBindRequest) SetChatIds(chatIds2 string) {
	this.ChatIds = chatIds2
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetChatIds() string {
	return this.ChatIds
}
func (this *OapiWorkspaceCorpGroupBindRequest) SetTargetCorpId(targetCorpId2 string) {
	this.TargetCorpId = targetCorpId2
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetTargetCorpId() string {
	return this.TargetCorpId
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.corp.group.bind"
}
func (this *OapiWorkspaceCorpGroupBindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCorpGroupBindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCorpGroupBindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_ids"] = this.ChatIds
	txtParams["target_corp_id"] = this.TargetCorpId
	return txtParams
}
func (this *OapiWorkspaceCorpGroupBindRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatIds, "chatIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCorpGroupBindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceCorpGroupBindResponse struct {
	taobao.TaobaoResponse
	Result []string `json:"result,omitempty"`
}
