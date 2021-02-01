package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHealthStepinfoGetuserstatusRequest() *OapiHealthStepinfoGetuserstatusRequest {
	return &OapiHealthStepinfoGetuserstatusRequest{}
}

type OapiHealthStepinfoGetuserstatusRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHealthStepinfoGetuserstatusResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiHealthStepinfoGetuserstatusRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHealthStepinfoGetuserstatusRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHealthStepinfoGetuserstatusRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHealthStepinfoGetuserstatusRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHealthStepinfoGetuserstatusRequest) GetApiMethodName() string {
	return "dingtalk.oapi.health.stepinfo.getuserstatus"
}
func (this *OapiHealthStepinfoGetuserstatusRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHealthStepinfoGetuserstatusRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHealthStepinfoGetuserstatusRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHealthStepinfoGetuserstatusRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHealthStepinfoGetuserstatusRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiHealthStepinfoGetuserstatusRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiHealthStepinfoGetuserstatusRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHealthStepinfoGetuserstatusRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHealthStepinfoGetuserstatusResponse struct {
	taobao.TaobaoResponse
	Status bool `json:"status,omitempty"`
}
