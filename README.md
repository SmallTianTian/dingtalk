# dingtalk

钉钉 Golang SDK

**注意：这不是钉钉官方 SDK，只是个人根据官方 JAVA SDK 翻译的。**

由于个人精力有限，无法将全部的接口验证通过，出现在 `test` 目录中的文件为已验证成功的，其余的文件，可自行根据官方 Java SDK 使用示例进行测试。

如果遇到问题，可提 [issue](https://github.com/SmallTianTian/dingtalk/issues/new)。如果能成功使用，也欢迎告知/提 `PR` 来完善整个测试。

### 使用方法和官方 `Java SDK`  基本一致。

#### 1. 泛型处理方式示例：

> 由于 Golang 不支持泛型，故 `client` 只返回网络错误，具体响应将放置在 `Request.Resp` 中。

**示例：**

[文档页面](https://ding-doc.dingtalk.com/document/app/obtain-application-suite-ticket)

请求示例（Java SDK）:

```java
DingTalkClient client = new DefaultDingTalkClient("https://oapi.dingtalk.com/service/get_suite_token");
OapiServiceGetSuiteTokenRequest req = new OapiServiceGetSuiteTokenRequest();
req.setSuiteKey("suitefcurkdvkc1nxxxx");
req.setSuiteSecret("y1ie2Rfb54xxxx");
req.setSuiteTicket("test");
OapiServiceGetSuiteTokenResponse rsp = client.execute(req);
System.out.println(rsp.getBody());
```

请求示例（Golang SDK）:

```go
// import "github.com/SmallTianTian/dingtalk/model"
// import "github.com/SmallTianTian/dingtalk"

client := dingtalk.NewDefaultDingTalkClient("https://oapi.dingtalk.com/service/get_suite_token")
req := model.NewOapiServiceGetSuiteTokenRequest()
req.SetSuiteKey("suitefcurkdvkc1nxxxx")
req.SetSuiteSecret("y1ie2Rfb54xxxx")
req.SetSuiteTicket("test")
err := client.Execute(req)
if err != nil {
  fmt.Println("Client error:", err)
  return
}

if req.Resp.IsSuccess() {
  fmt.Println("succ", req.Resp.Body)
} else {
  fmt.Println("fail", req.Resp.ErrMsg)
}
```

#### 2. 重载使用示例：

> 由于 Golang 不支持重载（重写），故 Golang SDK，相关部分需要改变。依据参数的个数，在原方法后面追加数字。
>
> 例如：原方法： `client.exec(one, two, three)`
>
> ​		    改变后：`client.Exec3(one, two, three)`

**示例：**

[文档页面](https://ding-doc.dingtalk.com/document/app/obtains-the-enterprise-authorized-credential)

请求示例（Java SDK）：

```java
DefaultDingTalkClient client = new DefaultDingTalkClient("https://oapi.dingtalk.com/service/get_corp_token");
OapiServiceGetCorpTokenRequest req = new OapiServiceGetCorpTokenRequest();
req.setAuthCorpid("dingc365fcabbf733c3535c2f4657eb6378f");
OapiServiceGetCorpTokenResponse execute = client.execute(req,"suiteKey","suiteSecrect", "suiteTicket");
```

请求示例（Golang SDK）:

```go
// import "github.com/SmallTianTian/dingtalk/model"
// import "github.com/SmallTianTian/dingtalk"

client := dingtalk.NewDefaultDingTalkClient("https://oapi.dingtalk.com/service/get_corp_token")
req := model.NewOapiServiceGetCorpTokenRequest()
req.SetAuthCorpid("dingc365fcabbf733c3535c2f4657eb6378f")
err := client.Execute4(req, "suiteKey", "suiteSecrect", "suiteTicket")
if err != nil {
  fmt.Println("Client error:", err)
  return
}

if req.Resp.IsSuccess() {
  fmt.Println("succ", req.Resp.Body)
} else {
  fmt.Println("fail", req.Resp.ErrMsg)
}
```

