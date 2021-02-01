package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRetailSellerQueryRequest() *OapiRetailSellerQueryRequest {
	return &OapiRetailSellerQueryRequest{}
}

type OapiRetailSellerQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRetailSellerQueryResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRetailSellerQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailSellerQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailSellerQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRetailSellerQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRetailSellerQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.seller.query"
}
func (this *OapiRetailSellerQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailSellerQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailSellerQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailSellerQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailSellerQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRetailSellerQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiRetailSellerQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailSellerQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRetailSellerQueryResponse struct {
	taobao.TaobaoResponse
	Result  []SellerDto `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
