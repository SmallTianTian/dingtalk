package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessBaseinfoListRequest() *OapiProcessBaseinfoListRequest {
	return &OapiProcessBaseinfoListRequest{}
}

type OapiProcessBaseinfoListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessBaseinfoListResponse
	ProcessCodes    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessBaseinfoListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessBaseinfoListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessBaseinfoListRequest) SetProcessCodes(processCodes2 string) {
	this.ProcessCodes = processCodes2
}
func (this *OapiProcessBaseinfoListRequest) GetProcessCodes() string {
	return this.ProcessCodes
}
func (this *OapiProcessBaseinfoListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.baseinfo.list"
}
func (this *OapiProcessBaseinfoListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessBaseinfoListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessBaseinfoListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessBaseinfoListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessBaseinfoListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_codes"] = this.ProcessCodes
	return txtParams
}
func (this *OapiProcessBaseinfoListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.ProcessCodes, 20, "processCodes"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessBaseinfoListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessBaseinfoListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessBaseinfoListResponse struct {
	taobao.TaobaoResponse
	Result  []ProcessTopVo `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type ProcessTopVo struct {
	BizCategoryId   string   `json:"biz_category_id,omitempty"`
	Description     string   `json:"description,omitempty"`
	ManagerUserIds  []string `json:"manager_user_ids,omitempty"`
	Name            string   `json:"name,omitempty"`
	OpenCustomPrint bool     `json:"open_custom_print,omitempty"`
	ProcessCode     string   `json:"process_code,omitempty"`
}
