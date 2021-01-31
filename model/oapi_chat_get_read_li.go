package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiChatGetReadListRequest() *OapiChatGetReadListRequest {
	return &OapiChatGetReadListRequest{}
}

type OapiChatGetReadListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatGetReadListResponse
	Cursor          int64
	MessageId       string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatGetReadListRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiChatGetReadListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatGetReadListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiChatGetReadListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiChatGetReadListRequest) SetMessageId(messageId2 string) {
	this.MessageId = messageId2
}
func (this *OapiChatGetReadListRequest) GetMessageId() string {
	return this.MessageId
}
func (this *OapiChatGetReadListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiChatGetReadListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiChatGetReadListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.getReadList"
}
func (this *OapiChatGetReadListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatGetReadListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatGetReadListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatGetReadListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatGetReadListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["messageId"] = this.MessageId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiChatGetReadListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiChatGetReadListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatGetReadListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatGetReadListResponse struct {
	taobao.TaobaoResponse
	Errcode        int64    `json:"errcode,omitempty"`
	Errmsg         string   `json:"errmsg,omitempty"`
	NextCursor     int64    `json:"next_cursor,omitempty"`
	ReadUserIdList []string `json:"readUserIdList,omitempty"`
}
