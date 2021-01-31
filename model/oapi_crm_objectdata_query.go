package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataQueryRequest() *OapiCrmObjectdataQueryRequest {
	return &OapiCrmObjectdataQueryRequest{}
}

type OapiCrmObjectdataQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataQueryResponse
	CurrentOperatorUserid string
	Cursor                string
	Name                  string
	PageSize              int64
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataQueryRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataQueryRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiCrmObjectdataQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiCrmObjectdataQueryRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiCrmObjectdataQueryRequest) GetName() string {
	return this.Name
}
func (this *OapiCrmObjectdataQueryRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiCrmObjectdataQueryRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiCrmObjectdataQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.query"
}
func (this *OapiCrmObjectdataQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["cursor"] = this.Cursor
	txtParams["name"] = this.Name
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiCrmObjectdataQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Name, "name"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataQueryResponse struct {
	taobao.TaobaoResponse
	Result IterablePage `json:"result,omitempty"`
}
