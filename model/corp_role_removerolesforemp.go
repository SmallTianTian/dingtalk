package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpRoleRemoverolesforempsRequest() *CorpRoleRemoverolesforempsRequest {
	return &CorpRoleRemoverolesforempsRequest{}
}

type CorpRoleRemoverolesforempsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpRoleRemoverolesforempsResponse
	RoleidList      string
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *CorpRoleRemoverolesforempsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpRoleRemoverolesforempsRequest) SetRoleidList(roleidList2 string) {
	this.RoleidList = roleidList2
}
func (this *CorpRoleRemoverolesforempsRequest) GetRoleidList() string {
	return this.RoleidList
}
func (this *CorpRoleRemoverolesforempsRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *CorpRoleRemoverolesforempsRequest) GetUseridList() string {
	return this.UseridList
}
func (this *CorpRoleRemoverolesforempsRequest) GetApiMethodName() string {
	return "dingtalk.corp.role.removerolesforemps"
}
func (this *CorpRoleRemoverolesforempsRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpRoleRemoverolesforempsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpRoleRemoverolesforempsRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpRoleRemoverolesforempsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpRoleRemoverolesforempsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpRoleRemoverolesforempsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["roleid_list"] = this.RoleidList
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *CorpRoleRemoverolesforempsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.RoleidList, "roleidList"); err != nil {
		return err
	}
	return nil
}
func (this *CorpRoleRemoverolesforempsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpRoleRemoverolesforempsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpRoleRemoverolesforempsResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
