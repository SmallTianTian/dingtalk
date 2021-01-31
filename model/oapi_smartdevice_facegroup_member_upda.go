package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartdeviceFacegroupMemberUpdateRequest() *OapiSmartdeviceFacegroupMemberUpdateRequest {
	return &OapiSmartdeviceFacegroupMemberUpdateRequest{}
}

type OapiSmartdeviceFacegroupMemberUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartdeviceFacegroupMemberUpdateResponse
	AddUserIds      string
	BizId           string
	DelUserIds      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) SetAddUserIds(addUserIds2 string) {
	this.AddUserIds = addUserIds2
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetAddUserIds() string {
	return this.AddUserIds
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetBizId() string {
	return this.BizId
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) SetDelUserIds(delUserIds2 string) {
	this.DelUserIds = delUserIds2
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetDelUserIds() string {
	return this.DelUserIds
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartdevice.facegroup.member.update"
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["add_user_ids"] = this.AddUserIds
	txtParams["biz_id"] = this.BizId
	txtParams["del_user_ids"] = this.DelUserIds
	return txtParams
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.AddUserIds, 100, "addUserIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartdeviceFacegroupMemberUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartdeviceFacegroupMemberUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
