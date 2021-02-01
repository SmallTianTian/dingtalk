package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiKefuSendmessageRequest() *OapiKefuSendmessageRequest {
	return &OapiKefuSendmessageRequest{}
}

type OapiKefuSendmessageRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKefuSendmessageResponse
	Content         string
	Customerid      string
	Msgtype         string
	Serviceid       int64
	Token           string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiKefuSendmessageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKefuSendmessageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKefuSendmessageRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiKefuSendmessageRequest) SetContentString(content2 string) {
	this.Content = content2
}
func (this *OapiKefuSendmessageRequest) GetContent() string {
	return this.Content
}
func (this *OapiKefuSendmessageRequest) SetCustomerid(customerid2 string) {
	this.Customerid = customerid2
}
func (this *OapiKefuSendmessageRequest) GetCustomerid() string {
	return this.Customerid
}
func (this *OapiKefuSendmessageRequest) SetMsgtype(msgtype2 string) {
	this.Msgtype = msgtype2
}
func (this *OapiKefuSendmessageRequest) GetMsgtype() string {
	return this.Msgtype
}
func (this *OapiKefuSendmessageRequest) SetServiceid(serviceid2 int64) {
	this.Serviceid = serviceid2
}
func (this *OapiKefuSendmessageRequest) GetServiceid() int64 {
	return this.Serviceid
}
func (this *OapiKefuSendmessageRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiKefuSendmessageRequest) GetToken() string {
	return this.Token
}
func (this *OapiKefuSendmessageRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiKefuSendmessageRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiKefuSendmessageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kefu.sendmessage"
}
func (this *OapiKefuSendmessageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKefuSendmessageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKefuSendmessageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKefuSendmessageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKefuSendmessageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["customerid"] = this.Customerid
	txtParams["msgtype"] = this.Msgtype
	txtParams["serviceid"] = this.Serviceid
	txtParams["token"] = this.Token
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiKefuSendmessageRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Content, taobao.DATA_CONTENT); err != nil {
		return err
	}
	return nil
}
func (this *OapiKefuSendmessageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKefuSendmessageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKefuSendmessageResponse struct {
	taobao.TaobaoResponse
	Taskid int64 `json:"taskid,omitempty"`
}
