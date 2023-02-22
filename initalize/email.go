package initalize

import (
	"gopkg.in/gomail.v2"
)

var Email = NewDefaultSendEmail()

// 阿里云默认屏蔽的25端口【开发465端口】
const (
	Host              = "smtp.qiye.aliyun.com"
	Port              = 465
	RustEmailUsername = "rust@dbsgw.cn"
	Password          = "ly062823.="
)

type RustSendEmail struct {
	Host     string
	Port     int
	Username string
	Password string
}

// NewDefaultSendEmail 默认配置发送邮件
func NewDefaultSendEmail() *RustSendEmail {
	return &RustSendEmail{
		Host:     Host,
		Port:     Port,
		Username: RustEmailUsername,
		Password: Password,
	}
}

// 发送验证码
func (e RustSendEmail) Send(mailTo string, body string) error {
	m := gomail.NewMessage()
	// 发件人信息
	m.SetHeader("From", m.FormatAddress(RustEmailUsername, "Rust中文网"))
	// 收件人
	m.SetHeader("To", mailTo)
	// 主题
	m.SetHeader("Subject", "Rust中文网")
	// 内容
	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	err := d.DialAndSend(m)
	if err != nil {
		// 处理错误
		return err
	}
	return nil
}

func init() {
	// 测试案例
	// if err := Email.Send("1578347363@qq.com", "Rust中文网", "<h1>来自Rust中文网验证码：654230222222</h1>"); err != nil {
	// 	panic("邮箱验证码发送失败")
	// } else {
	// 	fmt.Println("邮箱验证码发送成功")
	// }
}
