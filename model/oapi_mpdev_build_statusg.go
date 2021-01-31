package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMpdevBuildStatusGetRequest() *OapiMpdevBuildStatusGetRequest {
	return &OapiMpdevBuildStatusGetRequest{}
}

type OapiMpdevBuildStatusGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMpdevBuildStatusGetResponse
	MiniappId       string
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMpdevBuildStatusGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMpdevBuildStatusGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMpdevBuildStatusGetRequest) SetMiniappId(miniappId2 string) {
	this.MiniappId = miniappId2
}
func (this *OapiMpdevBuildStatusGetRequest) GetMiniappId() string {
	return this.MiniappId
}
func (this *OapiMpdevBuildStatusGetRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiMpdevBuildStatusGetRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiMpdevBuildStatusGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.mpdev.build.status.get"
}
func (this *OapiMpdevBuildStatusGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMpdevBuildStatusGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMpdevBuildStatusGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMpdevBuildStatusGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMpdevBuildStatusGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["miniapp_id"] = this.MiniappId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiMpdevBuildStatusGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MiniappId, "miniappId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMpdevBuildStatusGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMpdevBuildStatusGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMpdevBuildStatusGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  BuildResultVo `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
