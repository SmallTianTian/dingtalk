package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHrmEmployeeGetRequest() *OapiHrmEmployeeGetRequest {
	return &OapiHrmEmployeeGetRequest{}
}

type OapiHrmEmployeeGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHrmEmployeeGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiHrmEmployeeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHrmEmployeeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHrmEmployeeGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHrmEmployeeGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHrmEmployeeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hrm.employee.get"
}
func (this *OapiHrmEmployeeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHrmEmployeeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHrmEmployeeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHrmEmployeeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHrmEmployeeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiHrmEmployeeGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiHrmEmployeeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHrmEmployeeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHrmEmployeeGetResponse struct {
	taobao.TaobaoResponse
	Errcode   int64           `json:"errcode,omitempty"`
	Errmsg    string          `json:"errmsg,omitempty"`
	GroupList []FieldGroupVpo `json:"group_list,omitempty"`
}
