package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingCreateRequest() *OapiDingCreateRequest {
	return &OapiDingCreateRequest{}
}

type OapiDingCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingCreateResponse
	Attachment      string
	CreatorUserid   string
	ReceiverUserids string
	RemindTime      int64
	RemindType      int64
	TextContent     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingCreateRequest) SetAttachment(attachment2 string) {
	this.Attachment = attachment2
}
func (this *OapiDingCreateRequest) GetAttachment() string {
	return this.Attachment
}
func (this *OapiDingCreateRequest) SetCreatorUserid(creatorUserid2 string) {
	this.CreatorUserid = creatorUserid2
}
func (this *OapiDingCreateRequest) GetCreatorUserid() string {
	return this.CreatorUserid
}
func (this *OapiDingCreateRequest) SetReceiverUserids(receiverUserids2 string) {
	this.ReceiverUserids = receiverUserids2
}
func (this *OapiDingCreateRequest) GetReceiverUserids() string {
	return this.ReceiverUserids
}
func (this *OapiDingCreateRequest) SetRemindTime(remindTime2 int64) {
	this.RemindTime = remindTime2
}
func (this *OapiDingCreateRequest) GetRemindTime() int64 {
	return this.RemindTime
}
func (this *OapiDingCreateRequest) SetRemindType(remindType2 int64) {
	this.RemindType = remindType2
}
func (this *OapiDingCreateRequest) GetRemindType() int64 {
	return this.RemindType
}
func (this *OapiDingCreateRequest) SetTextContent(textContent2 string) {
	this.TextContent = textContent2
}
func (this *OapiDingCreateRequest) GetTextContent() string {
	return this.TextContent
}
func (this *OapiDingCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ding.create"
}
func (this *OapiDingCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attachment"] = this.Attachment
	txtParams["creator_userid"] = this.CreatorUserid
	txtParams["receiver_userids"] = this.ReceiverUserids
	txtParams["remind_time"] = this.RemindTime
	txtParams["remind_type"] = this.RemindType
	txtParams["text_content"] = this.TextContent
	return txtParams
}
func (this *OapiDingCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CreatorUserid, "creatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingCreateResponse struct {
	taobao.TaobaoResponse
	Result CorpDingCreateResult `json:"result,omitempty"`
}
