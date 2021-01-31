package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappScopeAddRequest() *OapiMicroappScopeAddRequest {
	return &OapiMicroappScopeAddRequest{}
}

type OapiMicroappScopeAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappScopeAddResponse
	Agentid         int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiMicroappScopeAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappScopeAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappScopeAddRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiMicroappScopeAddRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiMicroappScopeAddRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiMicroappScopeAddRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiMicroappScopeAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.scope.add"
}
func (this *OapiMicroappScopeAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappScopeAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappScopeAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappScopeAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappScopeAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiMicroappScopeAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappScopeAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappScopeAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappScopeAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
