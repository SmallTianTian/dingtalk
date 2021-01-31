package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHrmEmployeeUpdateresumerecordRequest() *OapiHrmEmployeeUpdateresumerecordRequest {
	return &OapiHrmEmployeeUpdateresumerecordRequest{}
}

type OapiHrmEmployeeUpdateresumerecordRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHrmEmployeeUpdateresumerecordResponse
	Content         string
	KVContent       string
	PcUrl           string
	PhoneUrl        string
	RecordTimestamp int64
	ResumeId        string
	Title           string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	WebUrl          string
}

func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetContent() string {
	return this.Content
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetkVContent(kVContent2 string) {
	this.KVContent = kVContent2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetkVContent() string {
	return this.KVContent
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetPcUrl(pcUrl2 string) {
	this.PcUrl = pcUrl2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetPcUrl() string {
	return this.PcUrl
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetPhoneUrl(phoneUrl2 string) {
	this.PhoneUrl = phoneUrl2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetPhoneUrl() string {
	return this.PhoneUrl
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetRecordTimestamp(recordTimestamp2 int64) {
	this.RecordTimestamp = recordTimestamp2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetRecordTimestamp() int64 {
	return this.RecordTimestamp
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetResumeId(resumeId2 string) {
	this.ResumeId = resumeId2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetResumeId() string {
	return this.ResumeId
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetTitle() string {
	return this.Title
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetWebUrl(webUrl2 string) {
	this.WebUrl = webUrl2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetWebUrl() string {
	return this.WebUrl
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hrm.employee.updateresumerecord"
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["k_v_content"] = this.KVContent
	txtParams["pc_url"] = this.PcUrl
	txtParams["phone_url"] = this.PhoneUrl
	txtParams["record_timestamp"] = this.RecordTimestamp
	txtParams["resume_id"] = this.ResumeId
	txtParams["title"] = this.Title
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["web_url"] = this.WebUrl
	return txtParams
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ResumeId, "resumeId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHrmEmployeeUpdateresumerecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHrmEmployeeUpdateresumerecordResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
