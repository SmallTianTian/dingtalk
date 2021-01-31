package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectMemberRemoveRequest() *OapiTdpProjectMemberRemoveRequest {
	return &OapiTdpProjectMemberRemoveRequest{}
}

type OapiTdpProjectMemberRemoveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectMemberRemoveResponse
	MemberId        string
	MicroappAgentId int64
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpProjectMemberRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectMemberRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectMemberRemoveRequest) SetMemberId(memberId2 string) {
	this.MemberId = memberId2
}
func (this *OapiTdpProjectMemberRemoveRequest) GetMemberId() string {
	return this.MemberId
}
func (this *OapiTdpProjectMemberRemoveRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectMemberRemoveRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectMemberRemoveRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectMemberRemoveRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectMemberRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.member.remove"
}
func (this *OapiTdpProjectMemberRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectMemberRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectMemberRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectMemberRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectMemberRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["member_id"] = this.MemberId
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	return txtParams
}
func (this *OapiTdpProjectMemberRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MemberId, "memberId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectMemberRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectMemberRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpProjectMemberRemoveResponse struct {
	taobao.TaobaoResponse
}
