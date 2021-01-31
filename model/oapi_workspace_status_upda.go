package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceStatusUpdateRequest() *OapiWorkspaceStatusUpdateRequest {
	return &OapiWorkspaceStatusUpdateRequest{}
}

type OapiWorkspaceStatusUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceStatusUpdateResponse
	TopHttpMethod   string
	TopResponseType string
	UpdateStatus    string
}

func (this *OapiWorkspaceStatusUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceStatusUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceStatusUpdateRequest) SetUpdateStatus(updateStatus2 string) {
	this.UpdateStatus = updateStatus2
}
func (this *OapiWorkspaceStatusUpdateRequest) GetUpdateStatus() string {
	return this.UpdateStatus
}
func (this *OapiWorkspaceStatusUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.status.update"
}
func (this *OapiWorkspaceStatusUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceStatusUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceStatusUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceStatusUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceStatusUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["update_status"] = this.UpdateStatus
	return txtParams
}
func (this *OapiWorkspaceStatusUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceStatusUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceStatusUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenWorkspaceUpdateStatusDto struct {
	Status string `json:"status,omitempty"`
}
type OapiWorkspaceStatusUpdateResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
