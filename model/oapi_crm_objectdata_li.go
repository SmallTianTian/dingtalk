package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataListRequest() *OapiCrmObjectdataListRequest {
	return &OapiCrmObjectdataListRequest{}
}

type OapiCrmObjectdataListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiCrmObjectdataListResponse
	CurrentOperatorUserid string
	DataIdList            string
	Name                  string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiCrmObjectdataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataListRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiCrmObjectdataListRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiCrmObjectdataListRequest) SetDataIdList(dataIdList2 string) {
	this.DataIdList = dataIdList2
}
func (this *OapiCrmObjectdataListRequest) GetDataIdList() string {
	return this.DataIdList
}
func (this *OapiCrmObjectdataListRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiCrmObjectdataListRequest) GetName() string {
	return this.Name
}
func (this *OapiCrmObjectdataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.list"
}
func (this *OapiCrmObjectdataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["data_id_list"] = this.DataIdList
	txtParams["name"] = this.Name
	return txtParams
}
func (this *OapiCrmObjectdataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DataIdList, "dataIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataListResponse struct {
	taobao.TaobaoResponse
	Errcode    int64                  `json:"errcode,omitempty"`
	Errmsg     string                 `json:"errmsg,omitempty"`
	ResultList []ObjectDataInstanceVo `json:"result_list,omitempty"`
}
