package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessTemplateListRequest() *OapiProcessTemplateListRequest {
	return &OapiProcessTemplateListRequest{}
}

type OapiProcessTemplateListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessTemplateListResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiProcessTemplateListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessTemplateListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessTemplateListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiProcessTemplateListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiProcessTemplateListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiProcessTemplateListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiProcessTemplateListRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProcessTemplateListRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProcessTemplateListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.template.list"
}
func (this *OapiProcessTemplateListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessTemplateListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessTemplateListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessTemplateListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessTemplateListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProcessTemplateListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessTemplateListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessTemplateListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessTemplateListResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  CorpTemplateVo `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type Templatelist struct {
	CanBeUpgraded bool   `json:"can_be_upgraded,omitempty"`
	CanModify     bool   `json:"can_modify,omitempty"`
	FormContent   string `json:"form_content,omitempty"`
	IconUrl       string `json:"icon_url,omitempty"`
	Name          string `json:"name,omitempty"`
	ProcessCode   string `json:"process_code,omitempty"`
	Url           string `json:"url,omitempty"`
}
type CorpTemplateVo struct {
	NextCursor   string         `json:"next_cursor,omitempty"`
	TemplateList []Templatelist `json:"template_list,omitempty"`
}
