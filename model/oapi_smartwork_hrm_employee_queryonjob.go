package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmEmployeeQueryonjobRequest() *OapiSmartworkHrmEmployeeQueryonjobRequest {
	return &OapiSmartworkHrmEmployeeQueryonjobRequest{}
}

type OapiSmartworkHrmEmployeeQueryonjobRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeQueryonjobResponse
	Offset          int64
	Size            int64
	StatusList      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) SetStatusList(statusList2 string) {
	this.StatusList = statusList2
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetStatusList() string {
	return this.StatusList
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.queryonjob"
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	txtParams["status_list"] = this.StatusList
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeQueryonjobRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeQueryonjobResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
