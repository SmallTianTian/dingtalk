package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAppstoreInternalRemindRequest() *OapiAppstoreInternalRemindRequest {
	return &OapiAppstoreInternalRemindRequest{}
}

type OapiAppstoreInternalRemindRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiAppstoreInternalRemindResponse
	GoodsCode         string
	ProcessInstanceId string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiAppstoreInternalRemindRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreInternalRemindRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreInternalRemindRequest) SetGoodsCode(goodsCode2 string) {
	this.GoodsCode = goodsCode2
}
func (this *OapiAppstoreInternalRemindRequest) GetGoodsCode() string {
	return this.GoodsCode
}
func (this *OapiAppstoreInternalRemindRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *OapiAppstoreInternalRemindRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *OapiAppstoreInternalRemindRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.internal.remind"
}
func (this *OapiAppstoreInternalRemindRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreInternalRemindRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreInternalRemindRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreInternalRemindRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreInternalRemindRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["goods_code"] = this.GoodsCode
	txtParams["process_instance_id"] = this.ProcessInstanceId
	return txtParams
}
func (this *OapiAppstoreInternalRemindRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAppstoreInternalRemindRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreInternalRemindRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreInternalRemindResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
