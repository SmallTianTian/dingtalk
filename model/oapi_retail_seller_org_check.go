package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRetailSellerOrgCheckRequest() *OapiRetailSellerOrgCheckRequest {
	return &OapiRetailSellerOrgCheckRequest{}
}

type OapiRetailSellerOrgCheckRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRetailSellerOrgCheckResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRetailSellerOrgCheckRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailSellerOrgCheckRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailSellerOrgCheckRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRetailSellerOrgCheckRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRetailSellerOrgCheckRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.seller.org.check"
}
func (this *OapiRetailSellerOrgCheckRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailSellerOrgCheckRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailSellerOrgCheckRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailSellerOrgCheckRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailSellerOrgCheckRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRetailSellerOrgCheckRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiRetailSellerOrgCheckRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailSellerOrgCheckRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRetailSellerOrgCheckResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
