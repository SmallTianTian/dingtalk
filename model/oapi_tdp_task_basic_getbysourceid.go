package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTaskBasicGetbysourceidRequest() *OapiTdpTaskBasicGetbysourceidRequest {
	return &OapiTdpTaskBasicGetbysourceidRequest{}
}

type OapiTdpTaskBasicGetbysourceidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTaskBasicGetbysourceidResponse
	AgentId         int64
	Source          string
	SourceId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpTaskBasicGetbysourceidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) SetSource(source2 string) {
	this.Source = source2
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetSource() string {
	return this.Source
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) SetSourceId(sourceId2 string) {
	this.SourceId = sourceId2
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetSourceId() string {
	return this.SourceId
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.task.basic.getbysourceid"
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["source"] = this.Source
	txtParams["source_id"] = this.SourceId
	return txtParams
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.SourceId, "sourceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTaskBasicGetbysourceidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpTaskBasicGetbysourceidResponse struct {
	taobao.TaobaoResponse
	Result OrgTask `json:"result,omitempty"`
}
