package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatSubadminUpdateRequest() *OapiChatSubadminUpdateRequest {
	return &OapiChatSubadminUpdateRequest{}
}

type OapiChatSubadminUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatSubadminUpdateResponse
	Chatid          string
	Role            int64
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiChatSubadminUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatSubadminUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatSubadminUpdateRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatSubadminUpdateRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatSubadminUpdateRequest) SetRole(role2 int64) {
	this.Role = role2
}
func (this *OapiChatSubadminUpdateRequest) GetRole() int64 {
	return this.Role
}
func (this *OapiChatSubadminUpdateRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiChatSubadminUpdateRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiChatSubadminUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.subadmin.update"
}
func (this *OapiChatSubadminUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatSubadminUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatSubadminUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatSubadminUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatSubadminUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatid"] = this.Chatid
	txtParams["role"] = this.Role
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiChatSubadminUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Chatid, "chatid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatSubadminUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatSubadminUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatSubadminUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
