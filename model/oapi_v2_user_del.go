package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2UserDeleteRequest() *OapiV2UserDeleteRequest {
	return &OapiV2UserDeleteRequest{}
}

type OapiV2UserDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2UserDeleteResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiV2UserDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2UserDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2UserDeleteRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2UserDeleteRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2UserDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.user.delete"
}
func (this *OapiV2UserDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2UserDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2UserDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2UserDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2UserDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiV2UserDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2UserDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2UserDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2UserDeleteResponse struct {
	taobao.TaobaoResponse
}
