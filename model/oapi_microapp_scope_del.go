package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappScopeDeleteRequest() *OapiMicroappScopeDeleteRequest {
	return &OapiMicroappScopeDeleteRequest{}
}

type OapiMicroappScopeDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappScopeDeleteResponse
	Agentid         int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiMicroappScopeDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappScopeDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappScopeDeleteRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiMicroappScopeDeleteRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiMicroappScopeDeleteRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiMicroappScopeDeleteRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiMicroappScopeDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.scope.delete"
}
func (this *OapiMicroappScopeDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappScopeDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappScopeDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappScopeDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappScopeDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiMicroappScopeDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappScopeDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappScopeDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappScopeDeleteResponse struct {
	taobao.TaobaoResponse
}
