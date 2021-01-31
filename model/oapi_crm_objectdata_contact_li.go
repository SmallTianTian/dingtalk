package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataContactListRequest() *OapiCrmObjectdataContactListRequest {
	return &OapiCrmObjectdataContactListRequest{}
}

type OapiCrmObjectdataContactListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataContactListResponse
	CurrentOperatorUserid string
	DataIdList            string
	ProviderCorpid        string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataContactListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataContactListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataContactListRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataContactListRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataContactListRequest) SetDataIdList(dataIdList2 string) {
	this.DataIdList = dataIdList2
}
func (this *OapiCrmObjectdataContactListRequest) GetDataIdList() string {
	return this.DataIdList
}
func (this *OapiCrmObjectdataContactListRequest) SetProviderCorpid(providerCorpid2 string) {
	this.ProviderCorpid = providerCorpid2
}
func (this *OapiCrmObjectdataContactListRequest) GetProviderCorpid() string {
	return this.ProviderCorpid
}
func (this *OapiCrmObjectdataContactListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.contact.list"
}
func (this *OapiCrmObjectdataContactListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataContactListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataContactListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataContactListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataContactListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["data_id_list"] = this.DataIdList
	txtParams["provider_corpid"] = this.ProviderCorpid
	return txtParams
}
func (this *OapiCrmObjectdataContactListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DataIdList, "dataIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataContactListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataContactListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataContactListResponse struct {
	taobao.TaobaoResponse
	Errcode    int64                  `json:"errcode,omitempty"`
	Errmsg     string                 `json:"errmsg,omitempty"`
	ResultList []ObjectDataInstanceVo `json:"result_list,omitempty"`
}
