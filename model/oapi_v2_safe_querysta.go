package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2SafeQuerystatusRequest() *OapiV2SafeQuerystatusRequest {
	return &OapiV2SafeQuerystatusRequest{}
}

type OapiV2SafeQuerystatusRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2SafeQuerystatusResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiV2SafeQuerystatusRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2SafeQuerystatusRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2SafeQuerystatusRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2SafeQuerystatusRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2SafeQuerystatusRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.safe.querystatus"
}
func (this *OapiV2SafeQuerystatusRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2SafeQuerystatusRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2SafeQuerystatusRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2SafeQuerystatusRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2SafeQuerystatusRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiV2SafeQuerystatusRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2SafeQuerystatusRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2SafeQuerystatusRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2SafeQuerystatusResponse struct {
	taobao.TaobaoResponse
	Result SafeQueryStatusResponse `json:"result,omitempty"`
}
type SafeQueryStatusResponse struct {
	Disable bool `json:"disable,omitempty"`
}
