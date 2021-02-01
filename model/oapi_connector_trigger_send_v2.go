package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiConnectorTriggerSendV2Request() *OapiConnectorTriggerSendV2Request {
	return &OapiConnectorTriggerSendV2Request{}
}

type OapiConnectorTriggerSendV2Request struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp              OapiConnectorTriggerSendV2Response
	TopHttpMethod     string
	TopResponseType   string
	TriggerMsgRequest string
}

func (this *OapiConnectorTriggerSendV2Request) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiConnectorTriggerSendV2Request) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiConnectorTriggerSendV2Request) SetTriggerMsgRequest(triggerMsgRequest2 string) {
	this.TriggerMsgRequest = triggerMsgRequest2
}
func (this *OapiConnectorTriggerSendV2Request) GetTriggerMsgRequest() string {
	return this.TriggerMsgRequest
}
func (this *OapiConnectorTriggerSendV2Request) GetApiMethodName() string {
	return "dingtalk.oapi.connector.trigger.send_v2"
}
func (this *OapiConnectorTriggerSendV2Request) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiConnectorTriggerSendV2Request) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiConnectorTriggerSendV2Request) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiConnectorTriggerSendV2Request) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiConnectorTriggerSendV2Request) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["trigger_msg_request"] = this.TriggerMsgRequest
	return txtParams
}
func (this *OapiConnectorTriggerSendV2Request) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiConnectorTriggerSendV2Request) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiConnectorTriggerSendV2Request) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TriggerMsgData struct {
	DdEventTime time.Time `json:"dd_event_time,omitempty"`
	JsonData    string    `json:"json_data,omitempty"`
	TriggerId   string    `json:"trigger_id,omitempty"`
}
type TriggerMsgRequest struct {
	Test               bool             `json:"test,omitempty"`
	TriggerMsgDataList []TriggerMsgData `json:"trigger_msg_data_list,omitempty"`
}
type OapiConnectorTriggerSendV2Response struct {
	taobao.TaobaoResponse
	Result TriggerMsgResponse `json:"result,omitempty"`
}
type TriggerMsgResponse struct {
	RequestId string `json:"request_id,omitempty"`
}
