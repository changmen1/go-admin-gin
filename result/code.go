// Package result 状态码和状态信息定义
// @author: zxl
package result

// Codes 定义状态
type Codes struct {
	Message map[uint]string
	Success uint
	Failed  uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	Success: 200,
	Failed:  501,
}

// 状态信息初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success: "成功",
		ApiCode.Failed:  "失败",
	}
}

// GetMessage 供外部调用

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
