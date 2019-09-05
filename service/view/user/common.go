package user

import (
	"encoding/json"
	"ginrest/service/config"
	"strconv"
	"time"

	"github.com/xxjwxc/oauth2/oauth2Client/src/data/view"
	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/myhttp"
)

//GetUserFromToken 通过token获取用户信息
func GetUserFromToken(token string) (username string, expireTime int, b bool) {
	if len(token) == 0 {
		return
	}

	//先从缓存中获取
	tmp, b := GetCacheBody(token)
	if b {
		username = tmp.User_name
		expireTime = tmp.Expire_time
		return
	}

	parm := make(map[string]string)
	parm["token"] = token
	bod, _ := json.Marshal(parm)
	rbody := myhttp.OnPostJSON(config.GetCheckTokenURL(), string(bod))
	if len(rbody) > 0 {
		var msg view.MapMessageBody
		json.Unmarshal([]byte(rbody), &msg)
		b = msg.State
		if msg.State {
			//保存缓存
			SaveCacheBody(token, msg.Data["username"], msg.Data["expire_time"])
			//返回结果
			username = msg.Data["username"]
			expireTime, _ = strconv.Atoi(msg.Data["expire_time"])
		}
	}

	return
}

//SaveCacheBody 保存缓存
func SaveCacheBody(accessToken, username, expireTime string) {
	var tmp view.UserCacheBody
	tmp.Access_token = accessToken
	tmp.User_name = username
	tmp.Expire_time, _ = strconv.Atoi(expireTime)
	//保存缓存
	cache := mycache.OnGetCache("oauth2")
	cache.Add(tmp.Access_token, &tmp, time.Duration(tmp.Expire_time)*time.Second)
	//------------------end
}

//GetCacheBody 获取缓存
func GetCacheBody(token string) (value *view.UserCacheBody, b bool) {
	cache := mycache.OnGetCache("oauth2")
	var tp interface{}
	tp, b = cache.Value(token)
	if b {
		value = tp.(*view.UserCacheBody)
	}

	return
}
