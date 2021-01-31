package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsChannelAccountDeleteRequest() *OapiAtsChannelAccountDeleteRequest {
	return &OapiAtsChannelAccountDeleteRequest{}
}

type OapiAtsChannelAccountDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiAtsChannelAccountDeleteResponse
	BizCode             string
	ChannelUserIdentify string
	TopHttpMethod       string
	TopResponseType     string
	Userid              string
}

func (this *OapiAtsChannelAccountDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsChannelAccountDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsChannelAccountDeleteRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsChannelAccountDeleteRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsChannelAccountDeleteRequest) SetChannelUserIdentify(channelUserIdentify2 string) {
	this.ChannelUserIdentify = channelUserIdentify2
}
func (this *OapiAtsChannelAccountDeleteRequest) GetChannelUserIdentify() string {
	return this.ChannelUserIdentify
}
func (this *OapiAtsChannelAccountDeleteRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAtsChannelAccountDeleteRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAtsChannelAccountDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.channel.account.delete"
}
func (this *OapiAtsChannelAccountDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsChannelAccountDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsChannelAccountDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsChannelAccountDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsChannelAccountDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["channel_user_identify"] = this.ChannelUserIdentify
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAtsChannelAccountDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsChannelAccountDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsChannelAccountDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsChannelAccountDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
}
