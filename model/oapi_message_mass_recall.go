package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageMassRecallRequest() *OapiMessageMassRecallRequest {
	return &OapiMessageMassRecallRequest{}
}

type OapiMessageMassRecallRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageMassRecallResponse
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiMessageMassRecallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageMassRecallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageMassRecallRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiMessageMassRecallRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiMessageMassRecallRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMessageMassRecallRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMessageMassRecallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.mass.recall"
}
func (this *OapiMessageMassRecallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageMassRecallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageMassRecallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageMassRecallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageMassRecallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["task_id"] = this.TaskId
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiMessageMassRecallRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskId, "taskId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageMassRecallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageMassRecallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageMassRecallResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
