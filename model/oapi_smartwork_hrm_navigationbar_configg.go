package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSmartworkHrmNavigationbarConfigGetRequest() *OapiSmartworkHrmNavigationbarConfigGetRequest {
	return &OapiSmartworkHrmNavigationbarConfigGetRequest{}
}

type OapiSmartworkHrmNavigationbarConfigGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmNavigationbarConfigGetResponse
	ChangeParam     string
	TopHttpMethod   string
	TopResponseType string
	Type            string
	Userid          string
}

func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) SetChangeParam(changeParam2 string) {
	this.ChangeParam = changeParam2
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetChangeParam() string {
	return this.ChangeParam
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetType() string {
	return this.Type
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.navigationbar.config.get"
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["change_param"] = this.ChangeParam
	txtParams["type"] = this.Type
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.ChangeParam, 20, "changeParam"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmNavigationbarConfigGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type HrmNavChangeVo struct {
	Code string `json:taobao.ERROR_COD,omitemptyE`
	Url  string `json:"url,omitempty"`
}
type OapiSmartworkHrmNavigationbarConfigGetResponse struct {
	taobao.TaobaoResponse
	Result  HrmNavigationBarVo `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type HrmNavBarTreePointLevel5Vo struct {
	AuthKey      string `json:"auth_key,omitempty"`
	AuthType     string `json:"auth_type,omitempty"`
	Code         string `json:taobao.ERROR_COD,omitemptyE`
	Icon         string `json:"icon,omitempty"`
	Name         string `json:"name,omitempty"`
	NoPermission bool   `json:"no_permission,omitempty"`
	Order        int64  `json:"order,omitempty"`
	Path         string `json:"path,omitempty"`
	SubAppCode   string `json:"sub_app_code,omitempty"`
	Url          string `json:"url,omitempty"`
}
type HrmNavBarTreePointLevel4Vo struct {
	AuthKey      string                       `json:"auth_key,omitempty"`
	AuthType     string                       `json:"auth_type,omitempty"`
	Children     []HrmNavBarTreePointLevel5Vo `json:"children,omitempty"`
	Code         string                       `json:taobao.ERROR_COD,omitemptyE`
	Icon         string                       `json:"icon,omitempty"`
	Name         string                       `json:"name,omitempty"`
	NoPermission bool                         `json:"no_permission,omitempty"`
	Order        int64                        `json:"order,omitempty"`
	Path         string                       `json:"path,omitempty"`
	SubAppCode   string                       `json:"sub_app_code,omitempty"`
	Url          string                       `json:"url,omitempty"`
}
type HrmNavBarTreePointLevel3Vo struct {
	AuthKey      string                       `json:"auth_key,omitempty"`
	AuthType     string                       `json:"auth_type,omitempty"`
	Children     []HrmNavBarTreePointLevel4Vo `json:"children,omitempty"`
	Code         string                       `json:taobao.ERROR_COD,omitemptyE`
	Icon         string                       `json:"icon,omitempty"`
	Name         string                       `json:"name,omitempty"`
	NoPermission bool                         `json:"no_permission,omitempty"`
	Order        int64                        `json:"order,omitempty"`
	Path         string                       `json:"path,omitempty"`
	SubAppCode   string                       `json:"sub_app_code,omitempty"`
	Url          string                       `json:"url,omitempty"`
}
type HrmNavBarTreePointLevel2Vo struct {
	AuthKey      string                       `json:"auth_key,omitempty"`
	AuthType     string                       `json:"auth_type,omitempty"`
	Children     []HrmNavBarTreePointLevel3Vo `json:"children,omitempty"`
	Code         string                       `json:taobao.ERROR_COD,omitemptyE`
	Icon         string                       `json:"icon,omitempty"`
	Name         string                       `json:"name,omitempty"`
	NoPermission bool                         `json:"no_permission,omitempty"`
	Order        int64                        `json:"order,omitempty"`
	Path         string                       `json:"path,omitempty"`
	SubAppCode   string                       `json:"sub_app_code,omitempty"`
	Url          string                       `json:"url,omitempty"`
}
type HrmNavBarTreePointLevel1Vo struct {
	AuthKey      string                       `json:"auth_key,omitempty"`
	AuthType     string                       `json:"auth_type,omitempty"`
	Children     []HrmNavBarTreePointLevel2Vo `json:"children,omitempty"`
	Code         string                       `json:taobao.ERROR_COD,omitemptyE`
	Icon         string                       `json:"icon,omitempty"`
	Name         string                       `json:"name,omitempty"`
	NoPermission bool                         `json:"no_permission,omitempty"`
	Order        int64                        `json:"order,omitempty"`
	Path         string                       `json:"path,omitempty"`
	SubAppCode   string                       `json:"sub_app_code,omitempty"`
	Url          string                       `json:"url,omitempty"`
}
type EmpBaseInfoVo struct {
	Avatar string `json:"avatar,omitempty"`
	Name   string `json:"name,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type HrmNavigationBarVo struct {
	CorpId      string                       `json:"corp_id,omitempty"`
	MenuList    []HrmNavBarTreePointLevel1Vo `json:"menu_list,omitempty"`
	SuperAdmins []EmpBaseInfoVo              `json:"super_admins,omitempty"`
	Userid      string                       `json:"userid,omitempty"`
}
