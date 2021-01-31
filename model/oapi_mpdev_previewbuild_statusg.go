package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMpdevPreviewbuildStatusGetRequest() *OapiMpdevPreviewbuildStatusGetRequest {
	return &OapiMpdevPreviewbuildStatusGetRequest{}
}

type OapiMpdevPreviewbuildStatusGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMpdevPreviewbuildStatusGetResponse
	MiniappId       string
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMpdevPreviewbuildStatusGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) SetMiniappId(miniappId2 string) {
	this.MiniappId = miniappId2
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetMiniappId() string {
	return this.MiniappId
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.mpdev.previewbuild.status.get"
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["miniapp_id"] = this.MiniappId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MiniappId, "miniappId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMpdevPreviewbuildStatusGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMpdevPreviewbuildStatusGetResponse struct {
	taobao.TaobaoResponse
	Result  BuildResultVo `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
