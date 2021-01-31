package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpHrmEmployeeAddresumerecordRequest() *CorpHrmEmployeeAddresumerecordRequest {
	return &CorpHrmEmployeeAddresumerecordRequest{}
}

type CorpHrmEmployeeAddresumerecordRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpHrmEmployeeAddresumerecordResponse
	Content         string
	KVContent       string
	PcUrl           string
	PhoneUrl        string
	RecordTimeStamp int64
	Title           string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WebUrl          string
}

func (this *CorpHrmEmployeeAddresumerecordRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetContent() string {
	return this.Content
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetkVContent(kVContent2 string) {
	this.KVContent = kVContent2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetkVContent() string {
	return this.KVContent
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetPcUrl(pcUrl2 string) {
	this.PcUrl = pcUrl2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetPcUrl() string {
	return this.PcUrl
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetPhoneUrl(phoneUrl2 string) {
	this.PhoneUrl = phoneUrl2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetPhoneUrl() string {
	return this.PhoneUrl
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetRecordTimeStamp(recordTimeStamp2 int64) {
	this.RecordTimeStamp = recordTimeStamp2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetRecordTimeStamp() int64 {
	return this.RecordTimeStamp
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetTitle() string {
	return this.Title
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetWebUrl(webUrl2 string) {
	this.WebUrl = webUrl2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetWebUrl() string {
	return this.WebUrl
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetApiMethodName() string {
	return "dingtalk.corp.hrm.employee.addresumerecord"
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpHrmEmployeeAddresumerecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["k_v_content"] = this.KVContent
	txtParams["pc_url"] = this.PcUrl
	txtParams["phone_url"] = this.PhoneUrl
	txtParams["record_time_stamp"] = this.RecordTimeStamp
	txtParams["title"] = this.Title
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["web_url"] = this.WebUrl
	return txtParams
}
func (this *CorpHrmEmployeeAddresumerecordRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpHrmEmployeeAddresumerecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpHrmEmployeeAddresumerecordResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
