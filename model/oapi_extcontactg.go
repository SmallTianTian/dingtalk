package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiExtcontactGetRequest() *OapiExtcontactGetRequest {
	return &OapiExtcontactGetRequest{}
}

type OapiExtcontactGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiExtcontactGetResponse
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *OapiExtcontactGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiExtcontactGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiExtcontactGetRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiExtcontactGetRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiExtcontactGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.extcontact.get"
}
func (this *OapiExtcontactGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiExtcontactGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiExtcontactGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiExtcontactGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiExtcontactGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["user_id"] = this.UserId
	return txtParams
}
func (this *OapiExtcontactGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UserId, "userId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiExtcontactGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiExtcontactGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiExtcontactGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  OpenExtContact `json:"result,omitempty"`
}
