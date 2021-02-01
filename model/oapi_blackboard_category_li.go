package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiBlackboardCategoryListRequest() *OapiBlackboardCategoryListRequest {
	return &OapiBlackboardCategoryListRequest{}
}

type OapiBlackboardCategoryListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBlackboardCategoryListResponse
	OperationUserid string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBlackboardCategoryListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBlackboardCategoryListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBlackboardCategoryListRequest) SetOperationUserid(operationUserid2 string) {
	this.OperationUserid = operationUserid2
}
func (this *OapiBlackboardCategoryListRequest) GetOperationUserid() string {
	return this.OperationUserid
}
func (this *OapiBlackboardCategoryListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.blackboard.category.list"
}
func (this *OapiBlackboardCategoryListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBlackboardCategoryListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBlackboardCategoryListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBlackboardCategoryListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBlackboardCategoryListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operation_userid"] = this.OperationUserid
	return txtParams
}
func (this *OapiBlackboardCategoryListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperationUserid, "operationUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiBlackboardCategoryListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBlackboardCategoryListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiBlackboardCategoryListResponse struct {
	taobao.TaobaoResponse
	Result  []BlackboardCategoryVo `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type BlackboardCategoryVo struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
