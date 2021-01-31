package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTasklistHiddenCancelRequest() *OapiTdpTasklistHiddenCancelRequest {
	return &OapiTdpTasklistHiddenCancelRequest{}
}

type OapiTdpTasklistHiddenCancelRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTasklistHiddenCancelResponse
	Agentid         int64
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiTdpTasklistHiddenCancelRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTasklistHiddenCancelRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiTdpTasklistHiddenCancelRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTasklistHiddenCancelRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.tasklist.hidden.cancel"
}
func (this *OapiTdpTasklistHiddenCancelRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTasklistHiddenCancelRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTasklistHiddenCancelRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiTdpTasklistHiddenCancelRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTasklistHiddenCancelRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTasklistHiddenCancelResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
