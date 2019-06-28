package user

// //获取全部关注用户
// func GetAll(w rest.ResponseWriter, r *rest.Request) {
// 	var req GetAllParam
// 	tools.GetRequestJsonObj(r, &req)
// 	//参数检测
// 	if !tools.CheckParam(req.Access_token) {
// 		w.WriteJson(message.GetErrorMsg(message.ParameterInvalid, req))
// 		return
// 	}
// 	//验证token，并获取用户名
// 	_, _, bfind := GetUserFromToken(req.Access_token)
// 	if !bfind {
// 		w.WriteJson(message.GetErrorMsg(message.TokenFailure))
// 		return
// 	}

// 	var db mysqldb.MySqlDB
// 	defer db.OnDestoryDB()
// 	orm := db.OnGetDBOrm(config.GetDbUrl())

// 	if req.Page_num == 0 {
// 		req.Page_num = 1 //默认第一页
// 	}
// 	var contentCount = config.GetContent_count() //每页10行数据
// 	var startCount = (req.Page_num - 1) * contentCount

// 	var users, all []*weixin.User_wx_info_tbl
// 	sql := orm.Where("headimg_url != ? ", "")
// 	sql2 := sql.Offset(startCount).Limit(contentCount)

// 	sql.Find(&all)
// 	sql2.Find(&users)

// 	if len(users) > 0 {
// 		for _, v := range users {
// 			v.Nick_name = common.UnicodeEmojiDecode(v.Nick_name)
// 		}

// 	}
// 	var total_page = tools.GetTotalPageNum(config.GetContent_count(), len(all)) //总页数
// 	mk := make(map[string]interface{})
// 	mk["total_page"] = total_page
// 	mk["total_count"] = len(all)
// 	mk["current_page"] = req.Page_num
// 	mk["current_data"] = users

// 	mess := message.GetSuccessMsg(message.NormalMessageId)
// 	mess.Data = mk
// 	w.WriteJson(mess)
// }
