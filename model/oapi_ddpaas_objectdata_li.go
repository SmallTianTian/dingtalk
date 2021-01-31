package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDdpaasObjectdataListRequest() *OapiDdpaasObjectdataListRequest {
	return &OapiDdpaasObjectdataListRequest{}
}

type OapiDdpaasObjectdataListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                  OapiDdpaasObjectdataListResponse
	AppUuid               string
	CurrentOperatorUserid string
	DataIdList            string
	FormCode              string
	TopHttpMethod         string
	TopResponseType       string
}

func (this *OapiDdpaasObjectdataListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDdpaasObjectdataListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDdpaasObjectdataListRequest) SetAppUuid(appUuid2 string) {
	this.AppUuid = appUuid2
}
func (this *OapiDdpaasObjectdataListRequest) GetAppUuid() string {
	return this.AppUuid
}
func (this *OapiDdpaasObjectdataListRequest) SetCurrentOperatorUserid(currentOperatorUserid2 string) {
	this.CurrentOperatorUserid = currentOperatorUserid2
}
func (this *OapiDdpaasObjectdataListRequest) GetCurrentOperatorUserid() string {
	return this.CurrentOperatorUserid
}
func (this *OapiDdpaasObjectdataListRequest) SetDataIdList(dataIdList2 string) {
	this.DataIdList = dataIdList2
}
func (this *OapiDdpaasObjectdataListRequest) GetDataIdList() string {
	return this.DataIdList
}
func (this *OapiDdpaasObjectdataListRequest) SetFormCode(formCode2 string) {
	this.FormCode = formCode2
}
func (this *OapiDdpaasObjectdataListRequest) GetFormCode() string {
	return this.FormCode
}
func (this *OapiDdpaasObjectdataListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ddpaas.objectdata.list"
}
func (this *OapiDdpaasObjectdataListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDdpaasObjectdataListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDdpaasObjectdataListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDdpaasObjectdataListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDdpaasObjectdataListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_uuid"] = this.AppUuid
	txtParams["current_operator_userid"] = this.CurrentOperatorUserid
	txtParams["data_id_list"] = this.DataIdList
	txtParams["form_code"] = this.FormCode
	return txtParams
}
func (this *OapiDdpaasObjectdataListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppUuid, "appUuid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDdpaasObjectdataListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDdpaasObjectdataListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDdpaasObjectdataListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                  `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Result  []ObjectDataInstanceVo `json:"result,omitempty"`
}
