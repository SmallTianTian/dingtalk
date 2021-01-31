package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCateringStoreGetRequest() *OapiCateringStoreGetRequest {
	return &OapiCateringStoreGetRequest{}
}

type OapiCateringStoreGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringStoreGetResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringStoreGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringStoreGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringStoreGetRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiCateringStoreGetRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiCateringStoreGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiCateringStoreGetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiCateringStoreGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.store.get"
}
func (this *OapiCateringStoreGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringStoreGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringStoreGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringStoreGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringStoreGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiCateringStoreGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCateringStoreGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringStoreGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringStoreGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type Managerlist struct {
	StaffId string `json:"staff_id,omitempty"`
}
type Data struct {
	DeptId      int64         `json:"dept_id,omitempty"`
	DeptName    string        `json:"dept_name,omitempty"`
	ManagerList []Managerlist `json:"manager_list,omitempty"`
	Type        string        `json:"type,omitempty"`
}
