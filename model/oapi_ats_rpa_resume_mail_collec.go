package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsRpaResumeMailCollectRequest() *OapiAtsRpaResumeMailCollectRequest {
	return &OapiAtsRpaResumeMailCollectRequest{}
}

type OapiAtsRpaResumeMailCollectRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsRpaResumeMailCollectResponse
	BizCode         string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsRpaResumeMailCollectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsRpaResumeMailCollectRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsRpaResumeMailCollectRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetParam() string {
	return this.Param
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.rpa.resume.mail.collect"
}
func (this *OapiAtsRpaResumeMailCollectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsRpaResumeMailCollectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsRpaResumeMailCollectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiAtsRpaResumeMailCollectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsRpaResumeMailCollectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type MailContent struct {
	BodyHtml        string `json:"body_html,omitempty"`
	FromMailAddress string `json:"from_mail_address,omitempty"`
	FromMailAlias   string `json:"from_mail_alias,omitempty"`
	MailAddress     string `json:"mail_address,omitempty"`
	MailId          string `json:"mail_id,omitempty"`
	ReceivedTime    int64  `json:"received_time,omitempty"`
	Title           string `json:"title,omitempty"`
}
type ResumeFileStoreVo struct {
	DownloadUrl string `json:"download_url,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	FileType    string `json:"file_type,omitempty"`
}
type OapiAtsRpaResumeMailCollectResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
