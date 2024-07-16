package main

import (
	"fmt"
	"testing"

	"github.com/ahuigo/requests"
)

const (
	// appID := viper.GetString("feishu.alarm.appID")
	// appSecret := viper.GetString("feishu.alarm.appSecret")
	FeishuTokenURL      = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	appID               = "xxx"
	appSecret           = "xxxx"
	FeishuDepartmentUrl = "https://open.feishu.cn/open-apis/contact/v3/departments"
	FeishuUserUrl       = "https://open.feishu.cn/open-apis/contact/v3/users"
)

type FeishuUser struct {
	ID     string `json:"userid" form:"userid"`
	Name   string `json:"name" form:"name"`
	EName  string `json:"ename" form:"ename"`
	Mobile string `json:"mobile" form:"mobile"`
	Email  string `json:"email" form:"email"`
}

type DepartmentRespData struct {
	HasMore   bool                 `json:"has_more"`
	PageToken string               `json:"page_token"`
	Items     []DepartmentRespItem `json:"items"`
}

type DepartmentResp struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data DepartmentRespData `json:"data"`
}

type UserResp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data UserRespData `json:"data"`
}
type UserRespData struct {
	HasMore   bool           `json:"has_more"`
	PageToken string         `json:"page_token"`
	Items     []UserRespItem `json:"items"`
}

type UserRespItem struct {
	OpenID string `json:"open_id"'`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}
type DepartmentRespItem struct {
	Name             string `json:"name"`
	DepartmentID     string `json:"department_id"`
	OpenDepartmentID string `json:"open_department_id"`
	Status           struct {
		IsDeleted bool `json:"is_deleted"`
	} `json:"status"`
}
type TokenResp struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int64  `json:"expire"`
}

var httpclient = requests.R().SetDebug()

func GetFeiShuToken(appID, appSecret string) (token string, err error) {
	params := requests.Json{
		"app_id":     appID,
		"app_secret": appSecret,
	}
	resp, err := httpclient.Post(FeishuTokenURL, params)
	if err != nil {
		return "", err
	}
	var tokenResp TokenResp
	err = resp.Json(&tokenResp)
	if err != nil {
		return "", fmt.Errorf("%s", err.Error())
	}
	if tokenResp.Code != 0 {
		return "", fmt.Errorf(" get token error,err:%s", tokenResp.Msg)
	}
	return tokenResp.TenantAccessToken, nil
}

var userlist []UserRespItem

func TestSyncDepartmentAndUserList(t *testing.T) {
	userlist = make([]UserRespItem, 0, 2000)
	// ctx := context.TODO()
	token, err := GetFeiShuToken(appID, appSecret)
	if err != nil {
		t.Fatal(err)
	}
	token = "Bearer " + token
	params := requests.Params{
		"parent_department_id": "0",
		"fetch_child":          "true",
		"page_size":            "50",
	}
	header := requests.Header{
		"Authorization": token,
	}
	for {
		// resp, err := requests.R().SetDebug().Get(FeishuDepartmentUrl, params, header)
		resp, err := httpclient.Get(FeishuDepartmentUrl, params, header)
		if err != nil {
			t.Logf("get token err,do http request error:%v", err)
			return
		}
		var departmentResp DepartmentResp
		err = resp.Json(&departmentResp)
		if err != nil {
			t.Logf(" unmarshal departmentResp err:%v", err)
			return
		}
		if departmentResp.Code != 0 {
			t.Logf(" get department err:%v,msg:%s,code:%d", err, departmentResp.Msg, departmentResp.Code)
			return
		}
		for _, item := range departmentResp.Data.Items {
			if !item.Status.IsDeleted {
				SyncUsers(t, &item, token)

			}
		}
		if !departmentResp.Data.HasMore {
			break
		}
		params["page_token"] = departmentResp.Data.PageToken
	}
	t.Log("userlist", userlist)
}

func SyncUsers(t *testing.T, item *DepartmentRespItem, token string) {
	// params := map[string]interface{}{}
	params := requests.Params{
		"department_id":      item.OpenDepartmentID,
		"page_size":          "100",
		"user_id_type":       "open_id",
		"department_id_type": "open_department_id",
	}

	for {
		// resp, err := requests.Get(FeishuUserUrl, params,
		resp, err := httpclient.Get(FeishuUserUrl, params,
			requests.Header{
				"Authorization": token,
			},
		)
		if err != nil {
			fmt.Printf("get user resp err:%v", err)
			return
		}
		var userResp UserResp
		err = resp.Json(&userResp)
		if err != nil {
			t.Logf(" unmarshal user resp err:%v", err)
			return
		}
		if userResp.Code != 0 {
			t.Logf("get userResp err msg:%s,code%d", userResp.Msg, userResp.Code)
			return
		}

		for _, item := range userResp.Data.Items {
			t.Log("name:", item.Name)
		}
		userlist = append(userlist, userResp.Data.Items...)

		if !userResp.Data.HasMore {
			break
		}
		params["page_token"] = userResp.Data.PageToken
	}
}
