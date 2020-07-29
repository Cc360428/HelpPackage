package utils

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"gopkg.in/gomail.v2"
)

func EmailCode(email string) (code string, err error) {
	logs.Info("输入的邮箱是：", email)
	code = GenValidateCode(6)
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "1097125645@qq.com", "Smart") // 发件人
	m.SetHeader("To", m.FormatAddress(email, "user"))        // 收件人
	m.SetHeader("Subject", "EmailCode")                      // 主题
	HTML := "<head>" +
		"<base target='_blank' />" +
		"<style type='text/css'>::-webkit-scrollbar{ display: none; }</style>" +
		"<style id='cloudAttachStyle' type='text/css'>#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>" +
		"<style id='blockquoteStyle' type='text/css'>blockquote{display:none;}</style>" +
		"<style type='text/css'>" +
		"	body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}" +
		"td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}" +
		"	pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%}" +
		"	th,td{font-family:arial,verdana,sans-serif;line-height:1.666}" +
		"	img{ border:0}" +
		"	header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}" +
		"	blockquote{margin-right:0px}" +
		"	</style>" +
		"	</head>" +
		"	<body tabindex='0'' role='listitem'>" +
		"	<table width='700' border='0' align='center' cellspacing='0' style='width:700px;'>" +
		"	<tbody>" +
		"	<tr>" +
		"	<td>" +
		"	<div style='width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;'>" +
		"	<table border='0' cellpadding='0' cellspacing='0' width='700'' height='39' style='font:12px Tahoma, Arial, 宋体;'>" +
		"	<tbody><tr><td width='210'></td></tr></tbody>" +
		"	</table>" +
		"	</div>" +
		"	<div style='width:680px;padding:0 10px;margin:0 auto;'>" +
		"	<div style='line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;'>" +
		"	<strong style='display:block;margin-bottom:15px;'>尊敬的用户： <span style='color:#f60;font-size: 16px;'></span>您好	!</strong>" +
		"	<strong style='display:block;margin-bottom:15px;'>" +
		"		</span>请在验证码输入框中输入：<span style='color:#5bddcf;font-size: 20px'>" + code + "</span>，完成操作" +
		"	</strong>" +
		"	</div>" +
		"	<div style='margin-bottom:30px;'>" +
		"	<small style='display:block;margin-bottom:20px;font-size:12px;'>" +
		"	<p style='color:#747474;'>" +
		"		注意：此操作可能会更改您的密码和登录邮箱。如果您不自行操作，请及时登录并更改密码，以确保帐户安全。" +
		"	<br>（工作人员不会要求您提供此验证码，请不要泄露！）" +
		"	</p>" +
		"	</small>" +
		"	</div>" +
		"	</div>" +
		"	<div style='width:700px;margin:0 auto;'>" +
		"	<div style='padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;'>" +
		"	<p>这是系统邮件，请不要回复<br>" +
		"		请妥善保管您的邮箱，以免他人盗用您的帐户" +
		"	</p>" +
		"	<p>有问题请联系：li_chao_cheng@163.com</p>" +
		"	</div>" +
		"	</div>" +
		"	</td>" +
		"	</tr>" +
		"	</tbody>" +
		"	</table>" +
		"	</body>"
	m.SetBody("text/html", HTML)
	d := gomail.NewDialer("smtp.qq.com", 465, "1097125645@qq.com", "xwmjrhrheymsiagh") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		logs.Error(err.Error())
		return code, err
	}
	return code, err
}

func Reply(email, title, content string) (err error) {
	logs.Info("输入的邮箱是：", email)
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "1097125645@qq.com", "li_chao_cheng.top") // 发件人
	m.SetHeader("To", m.FormatAddress(email, "user"))                    // 收件人
	m.SetHeader("Subject", "li_chao_cheng.top反馈")                        // 主题
	HTML := "<head>" +
		"<base target='_blank' />" +
		"<style type='text/css'>::-webkit-scrollbar{ display: none; }</style>" +
		"<style id='cloudAttachStyle' type='text/css'>#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style>" +
		"<style id='blockquoteStyle' type='text/css'>blockquote{display:none;}</style>" +
		"<style type='text/css'>" +
		"	body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}" +
		"td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}" +
		"	pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%}" +
		"	th,td{font-family:arial,verdana,sans-serif;line-height:1.666}" +
		"	img{ border:0}" +
		"	header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}" +
		"	blockquote{margin-right:0px}" +
		"	</style>" +
		"	</head>" +
		"	<body tabindex='0'' role='listitem'>" +
		"	<table width='700' border='0' align='center' cellspacing='0' style='width:700px;'>" +
		"	<tbody>" +
		"	<tr>" +
		"	<td>" +
		"	<div style='width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;'>" +
		"	<table border='0' cellpadding='0' cellspacing='0' width='700'' height='39' style='font:12px Tahoma, Arial, 宋体;'>" +
		"	<tbody><tr><td width='210'></td></tr></tbody>" +
		"	</table>" +
		"	</div>" +
		"	<div style='width:680px;padding:0 10px;margin:0 auto;'>" +
		"	<div style='line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;'>" +
		"	<strong style='display:block;margin-bottom:15px;'><span style='color:#f60;font-size: 16px;'></span>Hello!</strong>" +
		"	<strong style='display:block;margin-bottom:15px;'>" +
		"		</span>谢谢 您反馈<span style='color:#5bddcf;font-size: 20px'>" + title + " </span>" +
		"	</strong>" +
		"	</div>" +
		"	<div style='margin-bottom:30px;'>" +
		"	<small style='display:block;margin-bottom:20px;font-size:12px;'>" +
		"	<p style='color:#747474;'> 会根据你的反馈内容" + content + "会尽快完善" +
		"	<br>" +
		"	</p>" +
		"	</small>" +
		"	</div>" +
		"	</div>" +
		"	<div style='width:700px;margin:0 auto;'>" +
		"	<div style='padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;'>" +
		"	<p><br>" +
		"	</p>" +
		"	<p></p>" +
		"	</div>" +
		"	</div>" +
		"	</td>" +
		"	</tr>" +
		"	</tbody>" +
		"	</table>" +
		"	</body>"
	m.SetBody("text/html", HTML)
	d := gomail.NewDialer("smtp.qq.com", 465, "1097125645@qq.com", "xwmjrhrheymsiagh") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		logs.Error(err.Error())
		return err
	}
	return err
}
