package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingmiRobotGetRequest() *OapiDingmiRobotGetRequest {
	return &OapiDingmiRobotGetRequest{}
}

type OapiDingmiRobotGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingmiRobotGetResponse
	TopHttpMethod   string
	TopResponseType string
	Type            string
}

func (this *OapiDingmiRobotGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingmiRobotGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingmiRobotGetRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiDingmiRobotGetRequest) GetType() string {
	return this.Type
}
func (this *OapiDingmiRobotGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingmi.robot.get"
}
func (this *OapiDingmiRobotGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingmiRobotGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingmiRobotGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingmiRobotGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingmiRobotGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiDingmiRobotGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingmiRobotGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingmiRobotGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingmiRobotGetResponse struct {
	taobao.TaobaoResponse
	Result  CustomCrowdRobotDTO `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
type CustomCrowdRobotDTO struct {
	AppId           string `json:"app_id,omitempty"`
	Brief           string `json:"brief,omitempty"`
	ChatBoltId      string `json:"chat_bolt_id,omitempty"`
	Description     string `json:"description,omitempty"`
	IconUrl         string `json:"icon_url,omitempty"`
	Name            string `json:"name,omitempty"`
	PreviewMediaUrl string `json:"preview_media_url,omitempty"`
}
