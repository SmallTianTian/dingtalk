package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataCustomerListRequest() *OapiCrmObjectdataCustomerListRequest {
	return &OapiCrmObjectdataCustomerListRequest{}
}

type OapiCrmObjectdataCustomerListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataCustomerListResponse
	CurrentOperatorUserid string
	DataIdList            string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataCustomerListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataCustomerListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataCustomerListRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataCustomerListRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataCustomerListRequest) SetDataIdList(dataIdList2 string) {
	this.DataIdList = dataIdList2
}
func (this *OapiCrmObjectdataCustomerListRequest) GetDataIdList() string {
	return this.DataIdList
}
func (this *OapiCrmObjectdataCustomerListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.customer.list"
}
func (this *OapiCrmObjectdataCustomerListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataCustomerListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataCustomerListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataCustomerListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataCustomerListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["data_id_list"] = this.DataIdList
	return txtParams
}
func (this *OapiCrmObjectdataCustomerListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DataIdList, "dataIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataCustomerListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataCustomerListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataCustomerListResponse struct {
	taobao.TaobaoResponse

	ResultList []ObjectDataInstanceVo `json:"result_list,omitempty"`
}
