package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessTemplateManageGetRequest() *OapiProcessTemplateManageGetRequest {
	return &OapiProcessTemplateManageGetRequest{}
}

type OapiProcessTemplateManageGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessTemplateManageGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiProcessTemplateManageGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessTemplateManageGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessTemplateManageGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProcessTemplateManageGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProcessTemplateManageGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.template.manage.get"
}
func (this *OapiProcessTemplateManageGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessTemplateManageGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessTemplateManageGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessTemplateManageGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessTemplateManageGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProcessTemplateManageGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessTemplateManageGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessTemplateManageGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessTemplateManageGetResponse struct {
	taobao.TaobaoResponse
	Result  []ProcessSimpleVO `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type ProcessSimpleVO struct {
	AttendanceType int64     `json:"attendance_type,omitempty"`
	FlowTitle      string    `json:"flow_title,omitempty"`
	GmtModified    time.Time `json:"gmt_modified,omitempty"`
	IconName       string    `json:"icon_name,omitempty"`
	IconUrl        string    `json:"icon_url,omitempty"`
	IsNewProcess   bool      `json:"is_new_process,omitempty"`
	ProcessCode    string    `json:"process_code,omitempty"`
}
