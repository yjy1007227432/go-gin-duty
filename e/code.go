package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	QUESTTIMEOUT   = 405
	NOT_NIL_TOKEN  = 600

	ERROR_TIME_EARLY_FAIL = 10001

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_DECRYPT_TOKEN_FAIL       = 20005
	ERROR_NOT_ADMIN                = 20006
	ERROR_EXIST_AUTH               = 20007
	ERROR_ADD_AUTH_FAIL            = 20008
	ERROR_GET_NAME_FAIL            = 20009

	ERROR_EXIST_ROTA_FAIL        = 30001
	ERROR_NOT_EXIST_ROTA         = 30002
	ERROR_ADD_ROTA_FAIL          = 30003
	ERROR_IMPORT_ROTA_FAIL       = 30004
	ERROR_GET_ROTAS_FAIL         = 30005
	ERROR_DELETE_ROTAS_FAIL      = 30006
	ERROR_NOT_ROTAS_FAIL         = 30007
	ERROR_CHECK_ROTAS_EXIST_FAIL = 30008
	ERROR_EXIST_ROTA             = 30009
	NOT_NIL_MONTH                = 30010

	ERROR_TYPE_EXCEL     = 40001
	ERROR_BIND_DATA_FAIL = 40002

	ERROR_GET_RESTS_FAIL         = 50001
	ERROR_DELETE_RESTS_FAIL      = 50002
	ERROR_NOT_CHANGE_RESTS_FAIL  = 50003
	ERROR_EXAMINE_RESTS_FAIL     = 50004
	ERROR_NOT_EXAMINA_RESTS_FAIL = 50005
	ERROR_UPDATE_RESTS_FAIL      = 50006
	ERROR_REST_WEEKEND_FAIL      = 50007
	ERROR_ROTA_REST_FAIL         = 50008
	ERROR_EXIST_RESTS_FAIL       = 50009
	ERROR_ADD_RESTS_FAIL         = 50010

	ERROR_GET_AUTH_FAIL       = 60001
	ERROR_EXIST_USER_AUTH     = 60002
	ERROR_GENERATE_TOKEN_FAIL = 60003

	ERROR_GET_VACATION_FAIL    = 70001
	ERROR_DELETE_VACATION_FAIL = 70002
	ERROR_EDIT_VACATION_FAIL   = 70003

	ERROR_GET_EXCHANGE_FAIL           = 80001
	ERROR_NOT_CHANGE_EXCHANGE_FAIL    = 80002
	ERROR_DELETE_EXCHANGE_FAIL        = 80003
	ERROR_NOT_EXAMINA_EXCHANGE_FAIL   = 80004
	ERROR_UPDATE_EXCHANGE_FAIL        = 80005
	ERROR_RESPONCE_EXCHANGE_FAIL      = 80006
	ERROR_EXIST_EXCHANGE_FAIL         = 80007
	ERROR_ADD_EXCHANGE_FAIL           = 80008
	ERROR_EXCHANGE_SAME_FAIL          = 80009
	ERROR_DELETE_NOT_ME_EXCHANGE_FAIL = 80010
	ERROR_EXCHANGE_TYPE_FAIL          = 80011
	ERROR_REPLACE_DUTY_FAIL           = 80012

	ERROR_ADD_OVERTIME_FAIL         = 90001
	ERROR_GET_OVERTIME_FAIL         = 90002
	ERROR_NOT_EXAMINA_OVERTIME_FAIL = 90003
	ERROR_UPDATE_DUTYOVERTIME_FAIL  = 90004
	ERROR_DELETE_OVERTIME_FAIL      = 90005
)
