package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataContactQueryRequest() *OapiCrmObjectdataContactQueryRequest {
	return &OapiCrmObjectdataContactQueryRequest{}
}

type OapiCrmObjectdataContactQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataContactQueryResponse
	CurrentOperatorUserid string
	Cursor                string
	PageSize              int64
	ProviderCorpid        string
	QueryDsl              string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataContactQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataContactQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataContactQueryRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataContactQueryRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataContactQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiCrmObjectdataContactQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiCrmObjectdataContactQueryRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiCrmObjectdataContactQueryRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiCrmObjectdataContactQueryRequest) SetProviderCorpid(providerCorpid2 string) {
	this.ProviderCorpid = providerCorpid2
}
func (this *OapiCrmObjectdataContactQueryRequest) GetProviderCorpid() string {
	return this.ProviderCorpid
}
func (this *OapiCrmObjectdataContactQueryRequest) SetQueryDsl(queryDsl2 string) {
	this.QueryDsl = queryDsl2
}
func (this *OapiCrmObjectdataContactQueryRequest) GetQueryDsl() string {
	return this.QueryDsl
}
func (this *OapiCrmObjectdataContactQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.contact.query"
}
func (this *OapiCrmObjectdataContactQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataContactQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataContactQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataContactQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataContactQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["cursor"] = this.Cursor
	txtParams["page_size"] = this.PageSize
	txtParams["provider_corpid"] = this.ProviderCorpid
	txtParams["query_dsl"] = this.QueryDsl
	return txtParams
}
func (this *OapiCrmObjectdataContactQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PageSize, "pageSize"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataContactQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataContactQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataContactQueryResponse struct {
	taobao.TaobaoResponse
	Result IterablePage `json:"result,omitempty"`
}
type Values struct {
	CreatorNick    string           `json:"creator_nick,omitempty"`
	CreatorUserid  string           `json:"creator_userid,omitempty"`
	Data           string           `json:"data,omitempty"`
	ExtendData     string           `json:"extend_data,omitempty"`
	GmtCreate      string           `json:"gmt_create,omitempty"`
	GmtModified    string           `json:"gmt_modified,omitempty"`
	InstanceId     string           `json:"instance_id,omitempty"`
	ObjectType     string           `json:"object_type,omitempty"`
	Permission     DataPermissionVo `json:"permission,omitempty"`
	ProcInstStatus string           `json:"proc_inst_status,omitempty"`
	ProcOutResult  string           `json:"proc_out_result,omitempty"`
}
type IterablePage struct {
	HasMore    bool     `json:"has_more,omitempty"`
	NextCursor string   `json:"next_cursor,omitempty"`
	PageSize   int64    `json:"page_size,omitempty"`
	Values     []Values `json:"values,omitempty"`
}
