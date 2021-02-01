package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2DepartmentListparentbyuserRequest() *OapiV2DepartmentListparentbyuserRequest {
	return &OapiV2DepartmentListparentbyuserRequest{}
}

type OapiV2DepartmentListparentbyuserRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2DepartmentListparentbyuserResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiV2DepartmentListparentbyuserRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentListparentbyuserRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentListparentbyuserRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiV2DepartmentListparentbyuserRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiV2DepartmentListparentbyuserRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.listparentbyuser"
}
func (this *OapiV2DepartmentListparentbyuserRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentListparentbyuserRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentListparentbyuserRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentListparentbyuserRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentListparentbyuserRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiV2DepartmentListparentbyuserRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2DepartmentListparentbyuserRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentListparentbyuserRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentListparentbyuserResponse struct {
	taobao.TaobaoResponse
	Result DeptListParentByUserResponse `json:"result,omitempty"`
}
type DeptParentResponse struct {
	ParentDeptIdList []int64 `json:"parent_dept_id_list,omitempty"`
}
type DeptListParentByUserResponse struct {
	ParentList []DeptParentResponse `json:"parent_list,omitempty"`
}
