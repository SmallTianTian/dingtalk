package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImIntelligentCardSendRequest() *OapiImIntelligentCardSendRequest {
	return &OapiImIntelligentCardSendRequest{}
}

type OapiImIntelligentCardSendRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImIntelligentCardSendResponse
	BizType            int64
	OpenConversationId string
	TemplateData       string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImIntelligentCardSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImIntelligentCardSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImIntelligentCardSendRequest) SetBizType(bizType2 int64) {
	this.BizType = bizType2
}
func (this *OapiImIntelligentCardSendRequest) GetBizType() int64 {
	return this.BizType
}
func (this *OapiImIntelligentCardSendRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImIntelligentCardSendRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImIntelligentCardSendRequest) SetTemplateData(templateData2 string) {
	this.TemplateData = templateData2
}
func (this *OapiImIntelligentCardSendRequest) GetTemplateData() string {
	return this.TemplateData
}
func (this *OapiImIntelligentCardSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.intelligent.card.send"
}
func (this *OapiImIntelligentCardSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImIntelligentCardSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImIntelligentCardSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImIntelligentCardSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImIntelligentCardSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["template_data"] = this.TemplateData
	return txtParams
}
func (this *OapiImIntelligentCardSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizType, "bizType"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImIntelligentCardSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImIntelligentCardSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImIntelligentCardSendResponse struct {
	taobao.TaobaoResponse
	Result  SendInteractiveCardResultVo `json:"result,omitempty"`
	Success bool                        `json:"success,omitempty"`
}
type SendInteractiveCardResultVo struct {
	Result bool `json:"result,omitempty"`
}
