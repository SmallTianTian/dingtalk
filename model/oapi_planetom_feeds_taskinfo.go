package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiPlanetomFeedsTaskinfoRequest() *OapiPlanetomFeedsTaskinfoRequest {
	return &OapiPlanetomFeedsTaskinfoRequest{}
}

type OapiPlanetomFeedsTaskinfoRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPlanetomFeedsTaskinfoResponse
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPlanetomFeedsTaskinfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPlanetomFeedsTaskinfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPlanetomFeedsTaskinfoRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiPlanetomFeedsTaskinfoRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiPlanetomFeedsTaskinfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.planetom.feeds.taskinfo"
}
func (this *OapiPlanetomFeedsTaskinfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPlanetomFeedsTaskinfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPlanetomFeedsTaskinfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPlanetomFeedsTaskinfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPlanetomFeedsTaskinfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiPlanetomFeedsTaskinfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskId, "taskId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiPlanetomFeedsTaskinfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPlanetomFeedsTaskinfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPlanetomFeedsTaskinfoResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  []UploadFeedRspModel `json:"result,omitempty"`
}
type UploadFeedRspModel struct {
	ProcessMsg string `json:"process_msg,omitempty"`
	Success    bool   `json:"success,omitempty"`
	Title      string `json:"title,omitempty"`
}
