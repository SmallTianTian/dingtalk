package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduUserListRequest() *OapiEduUserListRequest {
	return &OapiEduUserListRequest{}
}

type OapiEduUserListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduUserListResponse
	ClassId         int64
	PageNo          int64
	PageSize        int64
	Role            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduUserListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduUserListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduUserListRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduUserListRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduUserListRequest) SetPageNo(pageNo2 int64) {
	this.PageNo = pageNo2
}
func (this *OapiEduUserListRequest) GetPageNo() int64 {
	return this.PageNo
}
func (this *OapiEduUserListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiEduUserListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiEduUserListRequest) SetRole(role2 string) {
	this.Role = role2
}
func (this *OapiEduUserListRequest) GetRole() string {
	return this.Role
}
func (this *OapiEduUserListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.user.list"
}
func (this *OapiEduUserListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduUserListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduUserListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduUserListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduUserListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["page_no"] = this.PageNo
	txtParams["page_size"] = this.PageSize
	txtParams["role"] = this.Role
	return txtParams
}
func (this *OapiEduUserListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduUserListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduUserListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduUserListResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type OpenEduUserDetail struct {
	ClassId int64  `json:"class_id,omitempty"`
	Feature string `json:"feature,omitempty"`
	Name    string `json:"name,omitempty"`
	Role    string `json:"role,omitempty"`
	Unionid string `json:"unionid,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
