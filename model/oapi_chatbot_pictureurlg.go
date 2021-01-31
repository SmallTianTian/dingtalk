package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatbotPictureurlGetRequest() *OapiChatbotPictureurlGetRequest {
	return &OapiChatbotPictureurlGetRequest{}
}

type OapiChatbotPictureurlGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatbotPictureurlGetResponse
	DownloadCode    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatbotPictureurlGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatbotPictureurlGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatbotPictureurlGetRequest) SetDownloadCode(downloadCode2 string) {
	this.DownloadCode = downloadCode2
}
func (this *OapiChatbotPictureurlGetRequest) GetDownloadCode() string {
	return this.DownloadCode
}
func (this *OapiChatbotPictureurlGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chatbot.pictureurl.get"
}
func (this *OapiChatbotPictureurlGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatbotPictureurlGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatbotPictureurlGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatbotPictureurlGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatbotPictureurlGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["download_code"] = this.DownloadCode
	return txtParams
}
func (this *OapiChatbotPictureurlGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DownloadCode, "downloadCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatbotPictureurlGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatbotPictureurlGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatbotPictureurlGetResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
