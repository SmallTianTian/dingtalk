package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOpenencryptEncryptboxStatusUpdateRequest() *OapiOpenencryptEncryptboxStatusUpdateRequest {
	return &OapiOpenencryptEncryptboxStatusUpdateRequest{}
}

type OapiOpenencryptEncryptboxStatusUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiOpenencryptEncryptboxStatusUpdateResponse
	TopEncryptBoxStatus string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) SetTopEncryptBoxStatus(topEncryptBoxStatus2 string) {
	this.TopEncryptBoxStatus = topEncryptBoxStatus2
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetTopEncryptBoxStatus() string {
	return this.TopEncryptBoxStatus
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.openencrypt.encryptbox.status.update"
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_encrypt_box_status"] = this.TopEncryptBoxStatus
	return txtParams
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOpenencryptEncryptboxStatusUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopEncryptBoxStatus struct {
	Appid     int64  `json:"appid,omitempty"`
	CorpId    string `json:"corp_id,omitempty"`
	Extension string `json:"extension,omitempty"`
	RequestId string `json:"request_id,omitempty"`
	Status    int64  `json:"status,omitempty"`
}
type OapiOpenencryptEncryptboxStatusUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
