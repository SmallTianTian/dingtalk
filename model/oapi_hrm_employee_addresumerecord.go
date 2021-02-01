package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiHrmEmployeeAddresumerecordRequest() *OapiHrmEmployeeAddresumerecordRequest {
	return &OapiHrmEmployeeAddresumerecordRequest{}
}

type OapiHrmEmployeeAddresumerecordRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHrmEmployeeAddresumerecordResponse
	Content         string
	KVContent       string
	PcUrl           string
	PhoneUrl        string
	RecordTimestamp int64
	Title           string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WebUrl          string
}

func (this *OapiHrmEmployeeAddresumerecordRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetContent() string {
	return this.Content
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetkVContent(kVContent2 string) {
	this.KVContent = kVContent2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetkVContent() string {
	return this.KVContent
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetPcUrl(pcUrl2 string) {
	this.PcUrl = pcUrl2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetPcUrl() string {
	return this.PcUrl
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetPhoneUrl(phoneUrl2 string) {
	this.PhoneUrl = phoneUrl2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetPhoneUrl() string {
	return this.PhoneUrl
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetRecordTimestamp(recordTimestamp2 int64) {
	this.RecordTimestamp = recordTimestamp2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetRecordTimestamp() int64 {
	return this.RecordTimestamp
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetTitle() string {
	return this.Title
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetWebUrl(webUrl2 string) {
	this.WebUrl = webUrl2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetWebUrl() string {
	return this.WebUrl
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hrm.employee.addresumerecord"
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHrmEmployeeAddresumerecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["k_v_content"] = this.KVContent
	txtParams["pc_url"] = this.PcUrl
	txtParams["phone_url"] = this.PhoneUrl
	txtParams["record_timestamp"] = this.RecordTimestamp
	txtParams["title"] = this.Title
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["web_url"] = this.WebUrl
	return txtParams
}
func (this *OapiHrmEmployeeAddresumerecordRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHrmEmployeeAddresumerecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHrmEmployeeAddresumerecordResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
