package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImpaasNewretailSendstaffmessageRequest() *OapiImpaasNewretailSendstaffmessageRequest {
	return &OapiImpaasNewretailSendstaffmessageRequest{}
}

type OapiImpaasNewretailSendstaffmessageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasNewretailSendstaffmessageResponse
	Content         string
	MsgType         int64
	Sender          string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiImpaasNewretailSendstaffmessageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetContentString(content2 string) {
	this.Content = content2
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetContent() string {
	return this.Content
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetMsgType(msgType2 int64) {
	this.MsgType = msgType2
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetMsgType() int64 {
	return this.MsgType
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetSender(sender2 string) {
	this.Sender = sender2
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetSender() string {
	return this.Sender
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.newretail.sendstaffmessage"
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["msg_type"] = this.MsgType
	txtParams["sender"] = this.Sender
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.UseridList, 100, "useridList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasNewretailSendstaffmessageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImpaasNewretailSendstaffmessageResponse struct {
	taobao.TaobaoResponse
	Result int64 `json:"result,omitempty"`
}
