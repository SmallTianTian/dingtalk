package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2SafeSetdisableRequest() *OapiV2SafeSetdisableRequest {
	return &OapiV2SafeSetdisableRequest{}
}

type OapiV2SafeSetdisableRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2SafeSetdisableResponse
	Reason          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiV2SafeSetdisableRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2SafeSetdisableRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2SafeSetdisableRequest) SetReason(reason2 string) {
	this.Reason = reason2
}
func (this *OapiV2SafeSetdisableRequest) GetReason() string {
	return this.Reason
}
func (this *OapiV2SafeSetdisableRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2SafeSetdisableRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2SafeSetdisableRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.safe.setdisable"
}
func (this *OapiV2SafeSetdisableRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2SafeSetdisableRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2SafeSetdisableRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2SafeSetdisableRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2SafeSetdisableRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["reason"] = this.Reason
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiV2SafeSetdisableRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Reason, "reason"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2SafeSetdisableRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2SafeSetdisableRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2SafeSetdisableResponse struct {
	taobao.TaobaoResponse
}
