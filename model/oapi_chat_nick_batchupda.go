package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatNickBatchupdateRequest() *OapiChatNickBatchupdateRequest {
	return &OapiChatNickBatchupdateRequest{}
}

type OapiChatNickBatchupdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatNickBatchupdateResponse
	Chatid          string
	TopHttpMethod   string
	TopResponseType string
	UserNickModel   string
}

func (this *OapiChatNickBatchupdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatNickBatchupdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatNickBatchupdateRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatNickBatchupdateRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatNickBatchupdateRequest) SetUserNickModel(userNickModel2 string) {
	this.UserNickModel = userNickModel2
}
func (this *OapiChatNickBatchupdateRequest) GetUserNickModel() string {
	return this.UserNickModel
}
func (this *OapiChatNickBatchupdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.nick.batchupdate"
}
func (this *OapiChatNickBatchupdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatNickBatchupdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatNickBatchupdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatNickBatchupdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatNickBatchupdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["user_nick_model"] = this.UserNickModel
	return txtParams
}
func (this *OapiChatNickBatchupdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatNickBatchupdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatNickBatchupdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenStaffIdAndNickModel struct {
	Nick   string `json:"nick,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type OapiChatNickBatchupdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
