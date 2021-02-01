package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpDingCreateRequest() *CorpDingCreateRequest {
	return &CorpDingCreateRequest{}
}

type CorpDingCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDingCreateResponse
	Attachment      string
	CreatorUserid   string
	ReceiverUserids string
	RemindTime      int64
	RemindType      int64
	TextContent     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDingCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDingCreateRequest) SetAttachment(attachment2 string) {
	this.Attachment = attachment2
}
func (this *CorpDingCreateRequest) GetAttachment() string {
	return this.Attachment
}
func (this *CorpDingCreateRequest) SetCreatorUserid(creatorUserid2 string) {
	this.CreatorUserid = creatorUserid2
}
func (this *CorpDingCreateRequest) GetCreatorUserid() string {
	return this.CreatorUserid
}
func (this *CorpDingCreateRequest) SetReceiverUserids(receiverUserids2 string) {
	this.ReceiverUserids = receiverUserids2
}
func (this *CorpDingCreateRequest) GetReceiverUserids() string {
	return this.ReceiverUserids
}
func (this *CorpDingCreateRequest) SetRemindTime(remindTime2 int64) {
	this.RemindTime = remindTime2
}
func (this *CorpDingCreateRequest) GetRemindTime() int64 {
	return this.RemindTime
}
func (this *CorpDingCreateRequest) SetRemindType(remindType2 int64) {
	this.RemindType = remindType2
}
func (this *CorpDingCreateRequest) GetRemindType() int64 {
	return this.RemindType
}
func (this *CorpDingCreateRequest) SetTextContent(textContent2 string) {
	this.TextContent = textContent2
}
func (this *CorpDingCreateRequest) GetTextContent() string {
	return this.TextContent
}
func (this *CorpDingCreateRequest) GetApiMethodName() string {
	return "dingtalk.corp.ding.create"
}
func (this *CorpDingCreateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDingCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDingCreateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDingCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDingCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDingCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attachment"] = this.Attachment
	txtParams["creator_userid"] = this.CreatorUserid
	txtParams["receiver_userids"] = this.ReceiverUserids
	txtParams["remind_time"] = this.RemindTime
	txtParams["remind_type"] = this.RemindType
	txtParams["text_content"] = this.TextContent
	return txtParams
}
func (this *CorpDingCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CreatorUserid, "creatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *CorpDingCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDingCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AttachmentVO struct {
	DetailType  string `json:"detail_type,omitempty"`
	FileId      string `json:"file_id,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	FileSize    int64  `json:"file_size,omitempty"`
	FileSpaceId string `json:"file_space_id,omitempty"`
	LinkPicUrl  string `json:"link_pic_url,omitempty"`
	LinkText    string `json:"link_text,omitempty"`
	LinkTitle   string `json:"link_title,omitempty"`
	LinkUrl     string `json:"link_url,omitempty"`
	Type        string `json:"type,omitempty"`
}
type CorpDingCreateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type CorpDingCreateResult struct {
	InvalidUsers []string `json:"invalid_users,omitempty"`
	OpenDingId   string   `json:"open_ding_id,omitempty"`
}
