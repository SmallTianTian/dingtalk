package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageCorpconversationSendbytemplateRequest() *OapiMessageCorpconversationSendbytemplateRequest {
	return &OapiMessageCorpconversationSendbytemplateRequest{}
}

type OapiMessageCorpconversationSendbytemplateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationSendbytemplateResponse
	AgentId         int64
	Data            string
	DeptIdList      string
	TemplateId      string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiMessageCorpconversationSendbytemplateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetData(data2 string) {
	this.Data = data2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetDataString(data2 string) {
	this.Data = data2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetData() string {
	return this.Data
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetDeptIdList(deptIdList2 string) {
	this.DeptIdList = deptIdList2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetDeptIdList() string {
	return this.DeptIdList
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetTemplateId(templateId2 string) {
	this.TemplateId = templateId2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetTemplateId() string {
	return this.TemplateId
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.sendbytemplate"
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["data"] = this.Data
	txtParams["dept_id_list"] = this.DeptIdList
	txtParams["template_id"] = this.TemplateId
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationSendbytemplateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageCorpconversationSendbytemplateResponse struct {
	taobao.TaobaoResponse
	TaskId int64 `json:"task_id,omitempty"`
}
