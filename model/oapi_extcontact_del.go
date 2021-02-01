package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiExtcontactDeleteRequest() *OapiExtcontactDeleteRequest {
	return &OapiExtcontactDeleteRequest{}
}

type OapiExtcontactDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiExtcontactDeleteResponse
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *OapiExtcontactDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiExtcontactDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiExtcontactDeleteRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiExtcontactDeleteRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiExtcontactDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.extcontact.delete"
}
func (this *OapiExtcontactDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiExtcontactDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiExtcontactDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiExtcontactDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiExtcontactDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["user_id"] = this.UserId
	return txtParams
}
func (this *OapiExtcontactDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UserId, "userId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiExtcontactDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiExtcontactDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiExtcontactDeleteResponse struct {
	taobao.TaobaoResponse
}
