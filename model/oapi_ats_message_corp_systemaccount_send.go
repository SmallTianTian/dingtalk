package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsMessageCorpSystemaccountSendRequest() *OapiAtsMessageCorpSystemaccountSendRequest {
	return &OapiAtsMessageCorpSystemaccountSendRequest{}
}

type OapiAtsMessageCorpSystemaccountSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsMessageCorpSystemaccountSendResponse
	BizCode         string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetParam() string {
	return this.Param
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.message.corp.systemaccount.send"
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsMessageCorpSystemaccountSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SystemAccountSendMessageParam struct {
	AccountKey         string   `json:"account_key,omitempty"`
	MessageKey         string   `json:"message_key,omitempty"`
	ReceiverUserIdList []string `json:"receiver_user_id_list,omitempty"`
	ValueMap           string   `json:"value_map,omitempty"`
}
type OapiAtsMessageCorpSystemaccountSendResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
