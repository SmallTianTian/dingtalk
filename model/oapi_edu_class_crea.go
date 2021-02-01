package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduClassCreateRequest() *OapiEduClassCreateRequest {
	return &OapiEduClassCreateRequest{}
}

type OapiEduClassCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduClassCreateResponse
	OpenClass       string
	Operator        string
	SuperId         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduClassCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduClassCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduClassCreateRequest) SetOpenClass(openClass2 string) {
	this.OpenClass = openClass2
}
func (this *OapiEduClassCreateRequest) GetOpenClass() string {
	return this.OpenClass
}
func (this *OapiEduClassCreateRequest) SetOperator(operator2 string) {
	this.Operator = operator2
}
func (this *OapiEduClassCreateRequest) GetOperator() string {
	return this.Operator
}
func (this *OapiEduClassCreateRequest) SetSuperId(superId2 int64) {
	this.SuperId = superId2
}
func (this *OapiEduClassCreateRequest) GetSuperId() int64 {
	return this.SuperId
}
func (this *OapiEduClassCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.class.create"
}
func (this *OapiEduClassCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduClassCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduClassCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduClassCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduClassCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_class"] = this.OpenClass
	txtParams["operator"] = this.Operator
	txtParams["super_id"] = this.SuperId
	return txtParams
}
func (this *OapiEduClassCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Operator, "operator"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduClassCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduClassCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenClass struct {
	ClassLevel  int64  `json:"class_level,omitempty"`
	Name        string `json:"name,omitempty"`
	Nick        string `json:"nick,omitempty"`
	OnlyUseNick string `json:"only_use_nick,omitempty"`
}
type OapiEduClassCreateResponse struct {
	taobao.TaobaoResponse
	Result  OpenClassCreateResponse `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type OpenClassCreateResponse struct {
	DeptId int64 `json:"dept_id,omitempty"`
}
