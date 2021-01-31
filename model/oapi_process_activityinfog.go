package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessActivityinfoGetRequest() *OapiProcessActivityinfoGetRequest {
	return &OapiProcessActivityinfoGetRequest{}
}

type OapiProcessActivityinfoGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessActivityinfoGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessActivityinfoGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessActivityinfoGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessActivityinfoGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessActivityinfoGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessActivityinfoGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.activityinfo.get"
}
func (this *OapiProcessActivityinfoGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessActivityinfoGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessActivityinfoGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessActivityinfoGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessActivityinfoGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessActivityinfoGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessActivityinfoGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessActivityinfoGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProcessQueryRequest struct {
	ProcessCode       string `json:"process_code,omitempty"`
	ProcessInstanceId string `json:"process_instance_id,omitempty"`
}
type OapiProcessActivityinfoGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  ProcessActivityTopVo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type ActivityModel struct {
	ActivityId string `json:"activity_id,omitempty"`
	Name       string `json:"name,omitempty"`
}
type ProcessActivityTopVo struct {
	Models      []ActivityModel `json:"models,omitempty"`
	ProcessCode string          `json:"process_code,omitempty"`
}
