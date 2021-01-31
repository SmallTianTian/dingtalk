package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessSyncRequest() *OapiProcessSyncRequest {
	return &OapiProcessSyncRequest{}
}

type OapiProcessSyncRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiProcessSyncResponse
	AgentId           int64
	BizCategoryId     string
	ProcessName       string
	SrcProcessCode    string
	TargetProcessCode string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiProcessSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessSyncRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiProcessSyncRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiProcessSyncRequest) SetBizCategoryId(bizCategoryId2 string) {
	this.BizCategoryId = bizCategoryId2
}
func (this *OapiProcessSyncRequest) GetBizCategoryId() string {
	return this.BizCategoryId
}
func (this *OapiProcessSyncRequest) SetProcessName(processName2 string) {
	this.ProcessName = processName2
}
func (this *OapiProcessSyncRequest) GetProcessName() string {
	return this.ProcessName
}
func (this *OapiProcessSyncRequest) SetSrcProcessCode(srcProcessCode2 string) {
	this.SrcProcessCode = srcProcessCode2
}
func (this *OapiProcessSyncRequest) GetSrcProcessCode() string {
	return this.SrcProcessCode
}
func (this *OapiProcessSyncRequest) SetTargetProcessCode(targetProcessCode2 string) {
	this.TargetProcessCode = targetProcessCode2
}
func (this *OapiProcessSyncRequest) GetTargetProcessCode() string {
	return this.TargetProcessCode
}
func (this *OapiProcessSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.sync"
}
func (this *OapiProcessSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["biz_category_id"] = this.BizCategoryId
	txtParams["process_name"] = this.ProcessName
	txtParams["src_process_code"] = this.SrcProcessCode
	txtParams["target_process_code"] = this.TargetProcessCode
	return txtParams
}
func (this *OapiProcessSyncRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessSyncResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
