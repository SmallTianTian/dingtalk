package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHireGuideBeginnertaskFinishRequest() *OapiHireGuideBeginnertaskFinishRequest {
	return &OapiHireGuideBeginnertaskFinishRequest{}
}

type OapiHireGuideBeginnertaskFinishRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHireGuideBeginnertaskFinishResponse
	TaskCode        string
	TaskType        int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiHireGuideBeginnertaskFinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHireGuideBeginnertaskFinishRequest) SetTaskCode(taskCode2 string) {
	this.TaskCode = taskCode2
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetTaskCode() string {
	return this.TaskCode
}
func (this *OapiHireGuideBeginnertaskFinishRequest) SetTaskType(taskType2 int64) {
	this.TaskType = taskType2
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetTaskType() int64 {
	return this.TaskType
}
func (this *OapiHireGuideBeginnertaskFinishRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.guide.beginnertask.finish"
}
func (this *OapiHireGuideBeginnertaskFinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHireGuideBeginnertaskFinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHireGuideBeginnertaskFinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["task_code"] = this.TaskCode
	txtParams["task_type"] = this.TaskType
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiHireGuideBeginnertaskFinishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskCode, "taskCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHireGuideBeginnertaskFinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHireGuideBeginnertaskFinishResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
