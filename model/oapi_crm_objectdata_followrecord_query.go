package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataFollowrecordQueryRequest() *OapiCrmObjectdataFollowrecordQueryRequest {
	return &OapiCrmObjectdataFollowrecordQueryRequest{}
}

type OapiCrmObjectdataFollowrecordQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataFollowrecordQueryResponse
	CurrentOperatorUserid string
	Cursor                string
	PageSize              int64
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.followrecord.query"
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["cursor"] = this.Cursor
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.PageSize, "pageSize"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataFollowrecordQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataFollowrecordQueryResponse struct {
	taobao.TaobaoResponse
	Result IterablePage `json:"result,omitempty"`
}
