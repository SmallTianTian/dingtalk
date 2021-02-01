package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCardIntelligentEmpgroupSendRequest() *OapiCardIntelligentEmpgroupSendRequest {
	return &OapiCardIntelligentEmpgroupSendRequest{}
}

type OapiCardIntelligentEmpgroupSendRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCardIntelligentEmpgroupSendResponse
	MsgKey          string
	ParamJson       string
	ReceiverList    string
	TopHttpMethod   string
	TopResponseType string
	Uuid            string
}

func (this *OapiCardIntelligentEmpgroupSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCardIntelligentEmpgroupSendRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiCardIntelligentEmpgroupSendRequest) SetParamJson(paramJson2 string) {
	this.ParamJson = paramJson2
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetParamJson() string {
	return this.ParamJson
}
func (this *OapiCardIntelligentEmpgroupSendRequest) SetReceiverList(receiverList2 string) {
	this.ReceiverList = receiverList2
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetReceiverList() string {
	return this.ReceiverList
}
func (this *OapiCardIntelligentEmpgroupSendRequest) SetUuid(uuid2 string) {
	this.Uuid = uuid2
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetUuid() string {
	return this.Uuid
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.card.intelligent.empgroup.send"
}
func (this *OapiCardIntelligentEmpgroupSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCardIntelligentEmpgroupSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCardIntelligentEmpgroupSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["msg_key"] = this.MsgKey
	txtParams["param_json"] = this.ParamJson
	txtParams["receiver_list"] = this.ReceiverList
	txtParams["uuid"] = this.Uuid
	return txtParams
}
func (this *OapiCardIntelligentEmpgroupSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MsgKey, "msgKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCardIntelligentEmpgroupSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCardIntelligentEmpgroupSendResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
