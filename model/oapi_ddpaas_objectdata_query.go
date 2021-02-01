package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDdpaasObjectdataQueryRequest() *OapiDdpaasObjectdataQueryRequest {
	return &OapiDdpaasObjectdataQueryRequest{}
}

type OapiDdpaasObjectdataQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiDdpaasObjectdataQueryResponse
	AppUuid               string
	CurrentOperatorUserid string
	Cursor                string
	FormCode              string
	QueryDsl              string
	Size                  int64
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiDdpaasObjectdataQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDdpaasObjectdataQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDdpaasObjectdataQueryRequest) SetAppUuid(appUuid2 string) {
	this.AppUuid = appUuid2
}
func (this *OapiDdpaasObjectdataQueryRequest) GetAppUuid() string {
	return this.AppUuid
}
func (this *OapiDdpaasObjectdataQueryRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiDdpaasObjectdataQueryRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiDdpaasObjectdataQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiDdpaasObjectdataQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiDdpaasObjectdataQueryRequest) SetFormCode(formCode2 string) {
	this.FormCode = formCode2
}
func (this *OapiDdpaasObjectdataQueryRequest) GetFormCode() string {
	return this.FormCode
}
func (this *OapiDdpaasObjectdataQueryRequest) SetQueryDsl(queryDsl2 string) {
	this.QueryDsl = queryDsl2
}
func (this *OapiDdpaasObjectdataQueryRequest) GetQueryDsl() string {
	return this.QueryDsl
}
func (this *OapiDdpaasObjectdataQueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiDdpaasObjectdataQueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiDdpaasObjectdataQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ddpaas.objectdata.query"
}
func (this *OapiDdpaasObjectdataQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDdpaasObjectdataQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDdpaasObjectdataQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDdpaasObjectdataQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDdpaasObjectdataQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_uuid"] = this.AppUuid
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["cursor"] = this.Cursor
	txtParams["form_code"] = this.FormCode
	txtParams["query_dsl"] = this.QueryDsl
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiDdpaasObjectdataQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppUuid, "appUuid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDdpaasObjectdataQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDdpaasObjectdataQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDdpaasObjectdataQueryResponse struct {
	taobao.TaobaoResponse
	Result IterablePage `json:"result,omitempty"`
}
