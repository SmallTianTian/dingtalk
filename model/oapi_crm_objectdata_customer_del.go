package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataCustomerDeleteRequest() *OapiCrmObjectdataCustomerDeleteRequest {
	return &OapiCrmObjectdataCustomerDeleteRequest{}
}

type OapiCrmObjectdataCustomerDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataCustomerDeleteResponse
	DataId          string
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataCustomerDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) SetDataId(dataId2 string) {
	this.DataId = dataId2
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetDataId() string {
	return this.DataId
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.customer.delete"
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["data_id"] = this.DataId
	txtParams["operator_userid"] = this.OperatorUserid
	return txtParams
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DataId, "dataId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataCustomerDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataCustomerDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64               `json:"errcode,omitempty"`
	Errmsg  string              `json:"errmsg,omitempty"`
	Result  ObjectDataDeleteDto `json:"result,omitempty"`
}
