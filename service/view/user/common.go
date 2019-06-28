package user

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/xxjwxc/dzmhotpot/service/config"
	"github.com/xxjwxc/oauth2/oauth2Client/src/data/view"
	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/myhttp"
)

//通过token获取用户信息
func GetUserFromToken(token string) (username string, expire_time int, b bool) {
	if len(token) == 0 {
		return
	}

	//先从缓存中获取
	tmp, b := GetCacheBody(token)
	if b {
		username = tmp.User_name
		expire_time = tmp.Expire_time
		return
	} else {
		parm := make(map[string]string)
		parm["token"] = token
		bod, _ := json.Marshal(parm)
		r_body := myhttp.OnPostJson(config.GetCheckTokenUrl(), string(bod))
		if len(r_body) == 0 {
			return
		} else {
			var msg view.MapMessageBody
			json.Unmarshal([]byte(r_body), &msg)
			b = msg.State
			if msg.State {
				//保存缓存
				SaveCacheBody(token, msg.Data["username"], msg.Data["expire_time"])

				//返回结果
				username = tmp.User_name
				expire_time = tmp.Expire_time
			}
			return
		}
	}
	//------------------end
}

//保存缓存
func SaveCacheBody(Access_token, username, Expire_time string) {
	var tmp view.UserCacheBody
	tmp.Access_token = Access_token
	tmp.User_name = username
	tmp.Expire_time, _ = strconv.Atoi(Expire_time)
	//保存缓存
	cache := mycache.OnGetCache("oauth2")
	cache.Add(tmp.Access_token, &tmp, time.Duration(tmp.Expire_time)*time.Second)
	//------------------end
}

//获取缓存
func GetCacheBody(token string) (value *view.UserCacheBody, b bool) {
	cache := mycache.OnGetCache("oauth2")
	var tp interface{}
	tp, b = cache.Value(token)
	if b {
		value = tp.(*view.UserCacheBody)
	}

	return
}
