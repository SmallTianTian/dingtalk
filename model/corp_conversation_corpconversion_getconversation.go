package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpConversationCorpconversionGetconversationRequest() *CorpConversationCorpconversionGetconversationRequest {
	return &CorpConversationCorpconversionGetconversationRequest{}
}

type CorpConversationCorpconversionGetconversationRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               CorpConversationCorpconversionGetconversationResponse
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *CorpConversationCorpconversionGetconversationRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpConversationCorpconversionGetconversationRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *CorpConversationCorpconversionGetconversationRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *CorpConversationCorpconversionGetconversationRequest) GetApiMethodName() string {
	return "dingtalk.corp.conversation.corpconversion.getconversation"
}
func (this *CorpConversationCorpconversionGetconversationRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpConversationCorpconversionGetconversationRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpConversationCorpconversionGetconversationRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpConversationCorpconversionGetconversationRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpConversationCorpconversionGetconversationRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpConversationCorpconversionGetconversationRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *CorpConversationCorpconversionGetconversationRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *CorpConversationCorpconversionGetconversationRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpConversationCorpconversionGetconversationRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpConversationCorpconversionGetconversationResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
