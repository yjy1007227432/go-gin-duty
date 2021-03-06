package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	QUESTTIMEOUT:   "超时",

	ERROR_TIME_EARLY_FAIL: "无法更新更早时间的信息",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_DECRYPT_TOKEN_FAIL:       "解析Token失败",
	ERROR_EXIST_AUTH:               "用户名密码已存在",
	ERROR_ADD_AUTH_FAIL:            "添加个人信息失败",
	ERROR_GET_NAME_FAIL:            "获取用户名失败",
	NOT_NIL_TOKEN:                  "Token 不能为空",

	ERROR_BIND_DATA_FAIL: "绑定数据失败失败",

	ERROR_GET_AUTH_FAIL:       "获取个人信息失败",
	ERROR_NOT_ADMIN:           "无管理员权限",
	ERROR_EXIST_USER_AUTH:     "相同用户已存在",
	ERROR_GENERATE_TOKEN_FAIL: "TOEKN解码失败",

	ERROR_EXIST_ROTA_FAIL:        "获取已存在值班日期失败",
	ERROR_EXIST_ROTA:             "已存在该日期值班情况",
	ERROR_NOT_EXIST_ROTA:         "不存在该日期值班情况",
	ERROR_ADD_ROTA_FAIL:          "插入日期值班情况失败",
	ERROR_IMPORT_ROTA_FAIL:       "导入值班表失败",
	ERROR_GET_ROTAS_FAIL:         "获取值班表失败",
	ERROR_DELETE_ROTAS_FAIL:      "删除当月值班表失败",
	ERROR_NOT_ROTAS_FAIL:         "非当天值班人员",
	ERROR_CHECK_ROTAS_EXIST_FAIL: "检测值班人员是否有值班失败",
	NOT_NIL_MONTH:                "参数不能为空",

	ERROR_TYPE_EXCEL: "excel表格标题格式错误",

	ERROR_GET_RESTS_FAIL:         "获取调休申请表失败",
	ERROR_DELETE_RESTS_FAIL:      "清空调休申请表失败",
	ERROR_NOT_CHANGE_RESTS_FAIL:  "已审批的调休申请无法删除",
	ERROR_NOT_EXAMINA_RESTS_FAIL: "已审批的调休申请无法重复审批",
	ERROR_EXAMINE_RESTS_FAIL:     "非审核人，调休申请无法审批",
	ERROR_UPDATE_RESTS_FAIL:      "更新调休申请表失败",
	ERROR_REST_WEEKEND_FAIL:      "周末调啥休？？？？",
	ERROR_ROTA_REST_FAIL:         "值班不允许调休",
	ERROR_EXIST_RESTS_FAIL:       "相同日期的调休表已存在",

	ERROR_GET_VACATION_FAIL:    "获取调休信息失败",
	ERROR_DELETE_VACATION_FAIL: "清空调休信息失败",
	ERROR_EDIT_VACATION_FAIL:   "修改调休信息失败",
	ERROR_ADD_RESTS_FAIL:       "新增调休信息失败",

	ERROR_GET_EXCHANGE_FAIL:           "获取换班申请表信息失败",
	ERROR_NOT_CHANGE_EXCHANGE_FAIL:    "已审批的换班申请无法删除",
	ERROR_DELETE_EXCHANGE_FAIL:        "删除调休申请失败",
	ERROR_NOT_EXAMINA_EXCHANGE_FAIL:   "已审批的换班申请无法重复审批",
	ERROR_UPDATE_EXCHANGE_FAIL:        "更新换班申请表失败",
	ERROR_RESPONCE_EXCHANGE_FAIL:      "非本人的换班申请无法审批",
	ERROR_DELETE_NOT_ME_EXCHANGE_FAIL: "非本人的换班申请无法删除",
	ERROR_EXIST_EXCHANGE_FAIL:         "已存在涉及相同日期的未处理换班请求表",
	ERROR_ADD_EXCHANGE_FAIL:           "新增换班请求失败",
	ERROR_EXCHANGE_SAME_FAIL:          "非同组人员或者本人与本人不能换班",
	ERROR_EXCHANGE_TYPE_FAIL:          "计费没有这么多换班类型",
	ERROR_REPLACE_DUTY_FAIL:           "顶班不需要填写被换班日期",

	ERROR_ADD_OVERTIME_FAIL:         "新增加班请求失败",
	ERROR_GET_OVERTIME_FAIL:         "获取加班申请表信息失败",
	ERROR_NOT_EXAMINA_OVERTIME_FAIL: "已审批的加班申请无法重复审批",
	ERROR_UPDATE_DUTYOVERTIME_FAIL:  "更新加班申请表失败",
	ERROR_DELETE_OVERTIME_FAIL:      "删除加班申请失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
