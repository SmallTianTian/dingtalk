package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiInactiveUserGetRequest() *OapiInactiveUserGetRequest {
	return &OapiInactiveUserGetRequest{}
}

type OapiInactiveUserGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiInactiveUserGetResponse
	Offset          int64
	QueryDate       string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiInactiveUserGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiInactiveUserGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiInactiveUserGetRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiInactiveUserGetRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiInactiveUserGetRequest) SetQueryDate(queryDate2 string) {
	this.QueryDate = queryDate2
}
func (this *OapiInactiveUserGetRequest) GetQueryDate() string {
	return this.QueryDate
}
func (this *OapiInactiveUserGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiInactiveUserGetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiInactiveUserGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.inactive.user.get"
}
func (this *OapiInactiveUserGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiInactiveUserGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiInactiveUserGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiInactiveUserGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiInactiveUserGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["query_date"] = this.QueryDate
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiInactiveUserGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiInactiveUserGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiInactiveUserGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiInactiveUserGetResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
