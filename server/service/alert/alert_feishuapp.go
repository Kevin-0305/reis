package alert

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type FSAPPConf struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type FSAPPTe struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type FSAPPElement struct {
	Tag           string         `json:"tag"`
	Text          Te             `json:"text"`
	Content       string         `json:"content"`
	FSAPPElements []FSAPPElement `json:"elements"`
}

type FSAPPTitles struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type FSAPPHeaders struct {
	FSAPPTitle FSAPPTitles `json:"title"`
	Template   string      `json:"template"`
}

type FSAPPCards struct {
	FSAPPConfig   FSAPPConf      `json:"config"`
	FSAPPElements []FSAPPElement `json:"elements"`
	FSAPPHeader   FSAPPHeaders   `json:"header"`
}

type FSAPP struct {
	MsgType       string     `json:"msg_type"`
	UnionIds      []string   `json:"union_ids"`      //@所使用字段 支持自定义部门ID，和open_department_id，列表长度小于等于 200 注：部门下的所有子部门包含的成员也会收到消息 示例值：["3dceba33a33226","d502aaa9514059", "od-5b91c9affb665451a16b90b4be367efa"]
	UserIds       []string   `json:"user_ids"`       //@所使用字段 用户 user_id 列表，长度小于等于 200 （对应 V3 接口的 employee_ids ） 示例值：["7cdcc7c2","ca51d83b"]
	OpenIds       []string   `json:"open_ids"`       //@所使用字段 用户 open_id 列表，长度小于等于 200 示例值：["ou_18eac85d35a26f989317ad4f02e8bbbb","ou_461cf042d9eedaa60d445f26dc747d5e"]
	DepartmentIds []string   `json:"department_ids"` //@所使用字段 用户 union_ids 列表，长度小于等于 200 示例值：["on_cad4860e7af114fb4ff6c5d496d1dd76","on_gdcq860e7af114fb4ff6c5d496dabcet"]
	FSAPPCard     FSAPPCards `json:"card"`
}

func GetAccessToken() (string, error) {
	// https://open.feishu.cn/open-apis/message/v4/batch_send/ 批量发送消息  tenant_access_token
	// 先获取 tenant_access_token
	u := TenantAccessMeg{
		AppId:     global.GVA_CONFIG.Alert.FsApp.AppID,
		AppSecret: global.GVA_CONFIG.Alert.FsApp.AppSecret,
	}
	fmt.Println("get token result", u)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	//var tr *http.Transport

	// if proxyUrl := beego.AppConfig.String("proxy"); proxyUrl != "" {
	// 	proxy := func(_ *http.Request) (*url.URL, error) {
	// 		return url.Parse(proxyUrl)
	// 	}
	// 	tr = &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 		Proxy:           proxy,
	// 	}
	// } else {
	// 	tr = &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 	}
	// }
	//client := &http.Client{Transport: tr}
	client := &http.Client{}
	//res, err := client.Post("https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal", "application/json; charset=utf-8", b)
	res, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal", b)
	if err != nil {
		global.GVA_LOG.Error("GetAccessToken error:", zap.Error(err))
		return "", err
	}
	res.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(res)
	defer res.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	fmt.Println("get token result", string(result))
	if err != nil {
		global.GVA_LOG.Error("GetAccessToken error:", zap.Error(err))
		return "", err
	}
	resp_json := TenantAccessResp{}
	json.Unmarshal(result, &resp_json)
	if resp_json.Msg != "ok" {
		global.GVA_LOG.Error("GetAccessToken error:", zap.Error(err))
		return "", errors.New(resp_json.Msg)
	}
	return resp_json.TenantAccessToken, nil
}

func PostToFeiShuApp(title, text, userIds string) string {
	open := "1"
	if open != "1" {
		global.GVA_LOG.Error("PostToFeiShuApp error:", zap.Error(errors.New("open-feishuapp is not open")))
		return ""
	}
	var color string
	if strings.Count(text, "resolved") > 0 && strings.Count(text, "firing") > 0 {
		color = "orange"
	} else if strings.Count(text, "resolved") > 0 {
		color = "green"
	} else {
		color = "red"
	}
	token, err := GetAccessToken()
	fmt.Println("Token", token)
	SendContent := text
	SendContentJson := []string{}
	if userIds != "" {
		UserIds := strings.Split(userIds, ",")
		UserIdtext := ""
		for _, UserId := range UserIds {
			UserIdtext += "<at user_id=" + UserId + "></at>"
			SendContentJson = append(SendContentJson, UserId)
		}

		SendContent += UserIdtext
	}

	u := FSAPP{
		MsgType:       "interactive",
		UnionIds:      SendContentJson,
		UserIds:       SendContentJson,
		OpenIds:       SendContentJson,
		DepartmentIds: SendContentJson,
		FSAPPCard: FSAPPCards{
			FSAPPConfig: FSAPPConf{
				WideScreenMode: true,
				EnableForward:  true,
			},
			FSAPPHeader: FSAPPHeaders{
				FSAPPTitle: FSAPPTitles{
					Content: title,
					Tag:     "plain_text",
				},
				Template: color,
			},
			FSAPPElements: []FSAPPElement{
				FSAPPElement{
					Tag: "div",
					Text: Te{
						Content: SendContent,
						Tag:     "lark_md",
					},
				},
				{
					Tag: "hr",
				},
				{
					Tag: "note",
					FSAPPElements: []FSAPPElement{
						{
							Content: title,
							Tag:     "lark_md",
						},
					},
				},
			},
		},
	}
	fmt.Println(u)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	global.GVA_LOG.Info("PostToFeiShuApp:", zap.String("SendContent", SendContent))
	//var tr *http.Transport
	// if proxyUrl := beego.AppConfig.String("proxy"); proxyUrl != "" {
	// 	proxy := func(_ *http.Request) (*url.URL, error) {
	// 		return url.Parse(proxyUrl)
	// 	}
	// 	tr = &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 		Proxy:           proxy,
	// 	}
	// } else {
	// 	tr = &http.Transport{
	// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// 	}
	// }
	//client := &http.Client{Transport: tr}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/message/v4/batch_send/", b)
	if err != nil {
		global.GVA_LOG.Error("PostToFeiShuApp error:", zap.Error(err))
		return ""
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("PostToFeiShuApp error:", zap.Error(err))
		return ""
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("PostToFeiShuApp error:", zap.Error(err))
		return ""
	}
	// resultMap := make(map[string]interface{})
	// json.Unmarshal(result, &resultMap)
	// if resultMap["msg"] != "ok" {
	// 	global.GVA_LOG.Error("PostToFeiShuApp error:", zap.Error(err))
	// 	return false, errors.New(resultMap["msg"].(string))
	// }
	// models.AlertToCounter.WithLabelValues("feishuapp").Add(1)
	// ChartsJson.Feishu += 1
	// logs.Info(logsign, "[feishuapp]", title+": "+string(result))
	return string(result)
}
