package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsChannelAccountAddRequest() *OapiAtsChannelAccountAddRequest {
	return &OapiAtsChannelAccountAddRequest{}
}

type OapiAtsChannelAccountAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiAtsChannelAccountAddResponse
	BizCode             string
	ChannelUserIdentify string
	TopHttpMethod       string
	TopResponseType     string
	Userid              string
}

func (this *OapiAtsChannelAccountAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsChannelAccountAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsChannelAccountAddRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsChannelAccountAddRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsChannelAccountAddRequest) SetChannelUserIdentify(channelUserIdentify2 string) {
	this.ChannelUserIdentify = channelUserIdentify2
}
func (this *OapiAtsChannelAccountAddRequest) GetChannelUserIdentify() string {
	return this.ChannelUserIdentify
}
func (this *OapiAtsChannelAccountAddRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAtsChannelAccountAddRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAtsChannelAccountAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.channel.account.add"
}
func (this *OapiAtsChannelAccountAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsChannelAccountAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsChannelAccountAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsChannelAccountAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsChannelAccountAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["channel_user_identify"] = this.ChannelUserIdentify
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAtsChannelAccountAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsChannelAccountAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsChannelAccountAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsChannelAccountAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
}
