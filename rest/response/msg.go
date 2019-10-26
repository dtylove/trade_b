package response

const OK = "0" // 成功

// common 10 00
const C_PARAMS_ERR = "1001"      // 参数错误
const C_TOKEN_NOT_FOUND = "1002" // token不存在

// user 11 00
const U_CREATE_USER_ERR = "1101" // 创建用户失败
const U_PWD_ERR = "1102"         // 密码错误
const U_EMAIL_NOT_FOUND = "1103" // 创建用户失败

// user 12 00
const O_ADD_ERR = "1201" // 创建订单失败
