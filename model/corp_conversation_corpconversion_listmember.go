package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpConversationCorpconversionListmemberRequest() *CorpConversationCorpconversionListmemberRequest {
	return &CorpConversationCorpconversionListmemberRequest{}
}

type CorpConversationCorpconversionListmemberRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               CorpConversationCorpconversionListmemberResponse
	Count              int64
	Offset             int64
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *CorpConversationCorpconversionListmemberRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpConversationCorpconversionListmemberRequest) SetCount(count2 int64) {
	this.Count = count2
}
func (this *CorpConversationCorpconversionListmemberRequest) GetCount() int64 {
	return this.Count
}
func (this *CorpConversationCorpconversionListmemberRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpConversationCorpconversionListmemberRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpConversationCorpconversionListmemberRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *CorpConversationCorpconversionListmemberRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *CorpConversationCorpconversionListmemberRequest) GetApiMethodName() string {
	return "dingtalk.corp.conversation.corpconversion.listmember"
}
func (this *CorpConversationCorpconversionListmemberRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpConversationCorpconversionListmemberRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpConversationCorpconversionListmemberRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpConversationCorpconversionListmemberRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpConversationCorpconversionListmemberRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpConversationCorpconversionListmemberRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["count"] = this.Count
	txtParams["offset"] = this.Offset
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *CorpConversationCorpconversionListmemberRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Count, "count"); err != nil {
		return err
	}
	return nil
}
func (this *CorpConversationCorpconversionListmemberRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpConversationCorpconversionListmemberRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpConversationCorpconversionListmemberResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
