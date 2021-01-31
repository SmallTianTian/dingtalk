package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiInactiveUserV2GetRequest() *OapiInactiveUserV2GetRequest {
	return &OapiInactiveUserV2GetRequest{}
}

type OapiInactiveUserV2GetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiInactiveUserV2GetResponse
	DeptIds         string
	IsActive        bool
	Offset          int64
	QueryDate       string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiInactiveUserV2GetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiInactiveUserV2GetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiInactiveUserV2GetRequest) SetDeptIds(deptIds2 string) {
	this.DeptIds = deptIds2
}
func (this *OapiInactiveUserV2GetRequest) GetDeptIds() string {
	return this.DeptIds
}
func (this *OapiInactiveUserV2GetRequest) SetIsActive(isActive2 bool) {
	this.IsActive = isActive2
}
func (this *OapiInactiveUserV2GetRequest) GetIsActive() bool {
	return this.IsActive
}
func (this *OapiInactiveUserV2GetRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiInactiveUserV2GetRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiInactiveUserV2GetRequest) SetQueryDate(queryDate2 string) {
	this.QueryDate = queryDate2
}
func (this *OapiInactiveUserV2GetRequest) GetQueryDate() string {
	return this.QueryDate
}
func (this *OapiInactiveUserV2GetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiInactiveUserV2GetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiInactiveUserV2GetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.inactive.user.v2.get"
}
func (this *OapiInactiveUserV2GetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiInactiveUserV2GetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiInactiveUserV2GetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiInactiveUserV2GetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiInactiveUserV2GetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_ids"] = this.DeptIds
	txtParams["is_active"] = this.IsActive
	txtParams["offset"] = this.Offset
	txtParams["query_date"] = this.QueryDate
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiInactiveUserV2GetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DeptIds, 100, "deptIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiInactiveUserV2GetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiInactiveUserV2GetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiInactiveUserV2GetResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
