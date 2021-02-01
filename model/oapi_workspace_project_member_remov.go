package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceProjectMemberRemoveRequest() *OapiWorkspaceProjectMemberRemoveRequest {
	return &OapiWorkspaceProjectMemberRemoveRequest{}
}

type OapiWorkspaceProjectMemberRemoveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceProjectMemberRemoveResponse
	OperatorStaffId string
	ProjectSourceId string
	Source          string
	StaffId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceProjectMemberRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) SetOperatorStaffId(operatorStaffId2 string) {
	this.OperatorStaffId = operatorStaffId2
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetOperatorStaffId() string {
	return this.OperatorStaffId
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) SetProjectSourceId(projectSourceId2 string) {
	this.ProjectSourceId = projectSourceId2
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetProjectSourceId() string {
	return this.ProjectSourceId
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) SetSource(source2 string) {
	this.Source = source2
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetSource() string {
	return this.Source
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) SetStaffId(staffId2 string) {
	this.StaffId = staffId2
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetStaffId() string {
	return this.StaffId
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.project.member.remove"
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operator_staff_id"] = this.OperatorStaffId
	txtParams["project_source_id"] = this.ProjectSourceId
	txtParams["source"] = this.Source
	txtParams["staff_id"] = this.StaffId
	return txtParams
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorStaffId, "operatorStaffId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceProjectMemberRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceProjectMemberRemoveResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
