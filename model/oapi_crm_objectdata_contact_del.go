package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectdataContactDeleteRequest() *OapiCrmObjectdataContactDeleteRequest {
	return &OapiCrmObjectdataContactDeleteRequest{}
}

type OapiCrmObjectdataContactDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataContactDeleteResponse
	DataId          string
	OperatorUserid  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataContactDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataContactDeleteRequest) SetDataId(dataId2 string) {
	this.DataId = dataId2
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetDataId() string {
	return this.DataId
}
func (this *OapiCrmObjectdataContactDeleteRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.contact.delete"
}
func (this *OapiCrmObjectdataContactDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataContactDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataContactDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["data_id"] = this.DataId
	txtParams["operator_userid"] = this.OperatorUserid
	return txtParams
}
func (this *OapiCrmObjectdataContactDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DataId, "dataId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataContactDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectdataContactDeleteResponse struct {
	taobao.TaobaoResponse
	Result ObjectDataDeleteDto `json:"result,omitempty"`
}
type ObjectDataDeleteDto struct {
	InstanceId string `json:"instance_id,omitempty"`
}
