package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessCopyRequest() *OapiProcessCopyRequest {
	return &OapiProcessCopyRequest{}
}

type OapiProcessCopyRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessCopyResponse
	AgentId         int64
	BizCategoryId   string
	CopyType        int64
	Description     string
	ProcessCode     string
	ProcessName     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessCopyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessCopyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessCopyRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiProcessCopyRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiProcessCopyRequest) SetBizCategoryId(bizCategoryId2 string) {
	this.BizCategoryId = bizCategoryId2
}
func (this *OapiProcessCopyRequest) GetBizCategoryId() string {
	return this.BizCategoryId
}
func (this *OapiProcessCopyRequest) SetCopyType(copyType2 int64) {
	this.CopyType = copyType2
}
func (this *OapiProcessCopyRequest) GetCopyType() int64 {
	return this.CopyType
}
func (this *OapiProcessCopyRequest) SetDescription(description2 string) {
	this.Description = description2
}
func (this *OapiProcessCopyRequest) GetDescription() string {
	return this.Description
}
func (this *OapiProcessCopyRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessCopyRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessCopyRequest) SetProcessName(processName2 string) {
	this.ProcessName = processName2
}
func (this *OapiProcessCopyRequest) GetProcessName() string {
	return this.ProcessName
}
func (this *OapiProcessCopyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.copy"
}
func (this *OapiProcessCopyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessCopyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessCopyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessCopyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessCopyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["biz_category_id"] = this.BizCategoryId
	txtParams["copy_type"] = this.CopyType
	txtParams["description"] = this.Description
	txtParams["process_code"] = this.ProcessCode
	txtParams["process_name"] = this.ProcessName
	return txtParams
}
func (this *OapiProcessCopyRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessCopyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessCopyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessCopyResponse struct {
	taobao.TaobaoResponse
	Result ProcessVo `json:"result,omitempty"`
}
type ProcessVo struct {
	BizCategoryId string `json:"biz_category_id,omitempty"`
	Description   string `json:"description,omitempty"`
	FormContent   string `json:"form_content,omitempty"`
	Name          string `json:"name,omitempty"`
	ProcessCode   string `json:"process_code,omitempty"`
}
