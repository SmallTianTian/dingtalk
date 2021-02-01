package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappCheckuidRequest() *OapiMicroappCheckuidRequest {
	return &OapiMicroappCheckuidRequest{}
}

type OapiMicroappCheckuidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappCheckuidResponse
	Agentid         int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiMicroappCheckuidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappCheckuidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappCheckuidRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiMicroappCheckuidRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiMicroappCheckuidRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiMicroappCheckuidRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiMicroappCheckuidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.checkuid"
}
func (this *OapiMicroappCheckuidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappCheckuidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappCheckuidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappCheckuidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappCheckuidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiMicroappCheckuidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappCheckuidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappCheckuidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappCheckuidResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
