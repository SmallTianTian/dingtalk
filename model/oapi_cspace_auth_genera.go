package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiCspaceAuthGenerateRequest() *OapiCspaceAuthGenerateRequest {
	return &OapiCspaceAuthGenerateRequest{}
}

type OapiCspaceAuthGenerateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceAuthGenerateResponse
	AgentId         int64
	AppId           int64
	Duration        int64
	FileIds         string
	Path            string
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *OapiCspaceAuthGenerateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCspaceAuthGenerateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceAuthGenerateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiCspaceAuthGenerateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiCspaceAuthGenerateRequest) SetAppId(appId2 int64) {
	this.AppId = appId2
}
func (this *OapiCspaceAuthGenerateRequest) GetAppId() int64 {
	return this.AppId
}
func (this *OapiCspaceAuthGenerateRequest) SetDuration(duration2 int64) {
	this.Duration = duration2
}
func (this *OapiCspaceAuthGenerateRequest) GetDuration() int64 {
	return this.Duration
}
func (this *OapiCspaceAuthGenerateRequest) SetFileIds(fileIds2 string) {
	this.FileIds = fileIds2
}
func (this *OapiCspaceAuthGenerateRequest) GetFileIds() string {
	return this.FileIds
}
func (this *OapiCspaceAuthGenerateRequest) SetPath(path2 string) {
	this.Path = path2
}
func (this *OapiCspaceAuthGenerateRequest) GetPath() string {
	return this.Path
}
func (this *OapiCspaceAuthGenerateRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiCspaceAuthGenerateRequest) GetType() string {
	return this.Type
}
func (this *OapiCspaceAuthGenerateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.auth.generate"
}
func (this *OapiCspaceAuthGenerateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceAuthGenerateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceAuthGenerateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceAuthGenerateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceAuthGenerateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["app_id"] = this.AppId
	txtParams["duration"] = this.Duration
	txtParams["file_ids"] = this.FileIds
	txtParams["path"] = this.Path
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiCspaceAuthGenerateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCspaceAuthGenerateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceAuthGenerateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceAuthGenerateResponse struct {
	taobao.TaobaoResponse
	Result  IsvAuthCodeResult `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type IsvAuthCodeResult struct {
	ExpireTime time.Time `json:"expire_time,omitempty"`
	IsvCode    string    `json:"isv_code,omitempty"`
}
