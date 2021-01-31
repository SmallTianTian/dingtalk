package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpRoleAddrolesforempsRequest() *CorpRoleAddrolesforempsRequest {
	return &CorpRoleAddrolesforempsRequest{}
}

type CorpRoleAddrolesforempsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpRoleAddrolesforempsResponse
	RolelidList     string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *CorpRoleAddrolesforempsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpRoleAddrolesforempsRequest) SetRolelidList(rolelidList2 string) {
	this.RolelidList = rolelidList2
}
func (this *CorpRoleAddrolesforempsRequest) GetRolelidList() string {
	return this.RolelidList
}
func (this *CorpRoleAddrolesforempsRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *CorpRoleAddrolesforempsRequest) GetUseridList() string {
	return this.UseridList
}
func (this *CorpRoleAddrolesforempsRequest) GetApiMethodName() string {
	return "dingtalk.corp.role.addrolesforemps"
}
func (this *CorpRoleAddrolesforempsRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpRoleAddrolesforempsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpRoleAddrolesforempsRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpRoleAddrolesforempsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpRoleAddrolesforempsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpRoleAddrolesforempsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rolelid_list"] = this.RolelidList
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *CorpRoleAddrolesforempsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RolelidList, "rolelidList"); err != nil {
		return err
	}
	return nil
}
func (this *CorpRoleAddrolesforempsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpRoleAddrolesforempsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpRoleAddrolesforempsResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
