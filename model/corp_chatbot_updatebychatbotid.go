package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpChatbotUpdatebychatbotidRequest() *CorpChatbotUpdatebychatbotidRequest {
	return &CorpChatbotUpdatebychatbotidRequest{}
}

type CorpChatbotUpdatebychatbotidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpChatbotUpdatebychatbotidResponse
	Breif           string
	ChatbotId       string
	Description     string
	Icon            string
	Name            string
	PreviewMediaId  string
	TopHttpMethod   string
	TopResponseType string
	UpdateType      int64
}

func (this *CorpChatbotUpdatebychatbotidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetBreif(breif2 string) {
	this.Breif = breif2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetBreif() string {
	return this.Breif
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetDescription(description2 string) {
	this.Description = description2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetDescription() string {
	return this.Description
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetIcon() string {
	return this.Icon
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetName() string {
	return this.Name
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetPreviewMediaId(previewMediaId2 string) {
	this.PreviewMediaId = previewMediaId2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetPreviewMediaId() string {
	return this.PreviewMediaId
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetUpdateType(updateType2 int64) {
	this.UpdateType = updateType2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetUpdateType() int64 {
	return this.UpdateType
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.updatebychatbotid"
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotUpdatebychatbotidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["breif"] = this.Breif
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["description"] = this.Description
	txtParams["icon"] = this.Icon
	txtParams["name"] = this.Name
	txtParams["preview_media_id"] = this.PreviewMediaId
	txtParams["update_type"] = this.UpdateType
	return txtParams
}
func (this *CorpChatbotUpdatebychatbotidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UpdateType, "updateType"); err != nil {
		return err
	}
	return nil
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotUpdatebychatbotidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpChatbotUpdatebychatbotidResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
