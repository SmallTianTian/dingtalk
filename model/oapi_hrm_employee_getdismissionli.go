package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHrmEmployeeGetdismissionlistRequest() *OapiHrmEmployeeGetdismissionlistRequest {
	return &OapiHrmEmployeeGetdismissionlistRequest{}
}

type OapiHrmEmployeeGetdismissionlistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHrmEmployeeGetdismissionlistResponse
	Current         int64
	OpUserid        string
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiHrmEmployeeGetdismissionlistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) SetCurrent(current2 int64) {
	this.Current = current2
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetCurrent() int64 {
	return this.Current
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hrm.employee.getdismissionlist"
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["current"] = this.Current
	txtParams["op_userid"] = this.OpUserid
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Current, "current"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHrmEmployeeGetdismissionlistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHrmEmployeeGetdismissionlistResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Page    HrmApiPage `json:"page,omitempty"`
}
