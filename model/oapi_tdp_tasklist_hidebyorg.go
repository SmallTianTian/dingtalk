package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTasklistHidebyorgRequest() *OapiTdpTasklistHidebyorgRequest {
	return &OapiTdpTasklistHidebyorgRequest{}
}

type OapiTdpTasklistHidebyorgRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTasklistHidebyorgResponse
	MicroappAgentId int64
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiTdpTasklistHidebyorgRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTasklistHidebyorgRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTasklistHidebyorgRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTasklistHidebyorgRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTasklistHidebyorgRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTasklistHidebyorgRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTasklistHidebyorgRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiTdpTasklistHidebyorgRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiTdpTasklistHidebyorgRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.tasklist.hidebyorg"
}
func (this *OapiTdpTasklistHidebyorgRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTasklistHidebyorgRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTasklistHidebyorgRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTasklistHidebyorgRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTasklistHidebyorgRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiTdpTasklistHidebyorgRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTasklistHidebyorgRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTasklistHidebyorgRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTasklistHidebyorgResponse struct {
	taobao.TaobaoResponse
}
