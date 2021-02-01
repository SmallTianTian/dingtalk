package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCallbackFailrecordConfirmRequest() *OapiCallbackFailrecordConfirmRequest {
	return &OapiCallbackFailrecordConfirmRequest{}
}

type OapiCallbackFailrecordConfirmRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallbackFailrecordConfirmResponse
	IdList          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCallbackFailrecordConfirmRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallbackFailrecordConfirmRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallbackFailrecordConfirmRequest) SetIdList(idList2 string) {
	this.IdList = idList2
}
func (this *OapiCallbackFailrecordConfirmRequest) GetIdList() string {
	return this.IdList
}
func (this *OapiCallbackFailrecordConfirmRequest) GetApiMethodName() string {
	return "dingtalk.oapi.callback.failrecord.confirm"
}
func (this *OapiCallbackFailrecordConfirmRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallbackFailrecordConfirmRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallbackFailrecordConfirmRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallbackFailrecordConfirmRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallbackFailrecordConfirmRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id_list"] = this.IdList
	return txtParams
}
func (this *OapiCallbackFailrecordConfirmRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.IdList, "idList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCallbackFailrecordConfirmRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallbackFailrecordConfirmRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallbackFailrecordConfirmResponse struct {
	taobao.TaobaoResponse
	ConfirmList []int64 `json:"confirm_list,omitempty"`
}
