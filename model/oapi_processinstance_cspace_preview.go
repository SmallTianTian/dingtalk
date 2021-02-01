package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessinstanceCspacePreviewRequest() *OapiProcessinstanceCspacePreviewRequest {
	return &OapiProcessinstanceCspacePreviewRequest{}
}

type OapiProcessinstanceCspacePreviewRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessinstanceCspacePreviewResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessinstanceCspacePreviewRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceCspacePreviewRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceCspacePreviewRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessinstanceCspacePreviewRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessinstanceCspacePreviewRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.cspace.preview"
}
func (this *OapiProcessinstanceCspacePreviewRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceCspacePreviewRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceCspacePreviewRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceCspacePreviewRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceCspacePreviewRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessinstanceCspacePreviewRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessinstanceCspacePreviewRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceCspacePreviewRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GrantCspaceRequest struct {
	Agentid           int64    `json:"agentid,omitempty"`
	FileId            string   `json:"file_id,omitempty"`
	FileidList        []string `json:"fileid_list,omitempty"`
	ProcessInstanceId string   `json:"process_instance_id,omitempty"`
	Userid            string   `json:"userid,omitempty"`
}
type OapiProcessinstanceCspacePreviewResponse struct {
	taobao.TaobaoResponse
	Result AppSpaceResponse `json:"result,omitempty"`
}
