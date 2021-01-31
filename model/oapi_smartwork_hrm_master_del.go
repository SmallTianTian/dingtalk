package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmMasterDeleteRequest() *OapiSmartworkHrmMasterDeleteRequest {
	return &OapiSmartworkHrmMasterDeleteRequest{}
}

type OapiSmartworkHrmMasterDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmMasterDeleteResponse
	BizData         string
	Tenantid        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmMasterDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmMasterDeleteRequest) SetBizData(bizData2 string) {
	this.BizData = bizData2
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetBizData() string {
	return this.BizData
}
func (this *OapiSmartworkHrmMasterDeleteRequest) SetTenantid(tenantid2 int64) {
	this.Tenantid = tenantid2
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetTenantid() int64 {
	return this.Tenantid
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.master.delete"
}
func (this *OapiSmartworkHrmMasterDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmMasterDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmMasterDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_data"] = this.BizData
	txtParams["tenantid"] = this.Tenantid
	return txtParams
}
func (this *OapiSmartworkHrmMasterDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.BizData, 20, "bizData"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmMasterDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Fieldlist struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type ScopeVo struct {
	ScopeCode string `json:"scope_code,omitempty"`
	Version   int64  `json:"version,omitempty"`
}
type BizDataVo struct {
	BizTime int64       `json:"biz_time,omitempty"`
	BizUk   string      `json:"biz_uk,omitempty"`
	Fields  []Fieldlist `json:"fields,omitempty"`
	Scope   ScopeVo     `json:"scope,omitempty"`
	Userid  string      `json:"userid,omitempty"`
}
type OapiSmartworkHrmMasterDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64                  `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Result  IntegrateBatchResultVo `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
