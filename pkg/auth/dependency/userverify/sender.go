package userverify

import (
	"errors"
	"fmt"

	"github.com/skygeario/skygear-server/pkg/core/sms"

	"github.com/go-gomail/gomail"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userprofile"
	authTemplate "github.com/skygeario/skygear-server/pkg/auth/template"
	"github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/mail"
	"github.com/skygeario/skygear-server/pkg/core/template"
)

type CodeSender interface {
	Send(verifyCode VerifyCode, userProfile userprofile.UserProfile) error
}

type EmailCodeSender struct {
	AppName        string
	Config         config.UserVerifyConfiguration
	Dialer         *gomail.Dialer
	TemplateEngine *template.Engine
}

func (e *EmailCodeSender) Send(verifyCode VerifyCode, userProfile userprofile.UserProfile) (err error) {
	var keyConfig config.UserVerifyKeyConfiguration
	var ok bool
	if keyConfig, ok = e.Config.ConfigForKey(verifyCode.RecordKey); !ok {
		return errors.New("provider for " + verifyCode.RecordKey + " not found")
	}

	context := prepareVerifyRequestContext(
		verifyCode,
		e.AppName,
		e.Config,
		userProfile,
	)

	providerConfig := keyConfig.ProviderConfig

	var textBody string
	if textBody, err = e.TemplateEngine.ParseTextTemplate(
		authTemplate.VerifyTextTemplateNameForKey(verifyCode.RecordKey),
		context,
		template.ParseOption{Required: true, FallbackTemplateName: authTemplate.TemplateNameVerifyEmailText},
	); err != nil {
		return
	}

	var htmlBody string
	if htmlBody, err = e.TemplateEngine.ParseTextTemplate(
		authTemplate.VerifyHTMLTemplateNameForKey(verifyCode.RecordKey),
		context,
		template.ParseOption{Required: false, FallbackTemplateName: authTemplate.TemplateNameVerifyEmailHTML},
	); err != nil {
		return
	}

	sendReq := mail.SendRequest{
		Dialer:      e.Dialer,
		Sender:      providerConfig.Sender,
		SenderName:  providerConfig.SenderName,
		Recipient:   verifyCode.RecordValue,
		Subject:     providerConfig.Subject,
		ReplyTo:     providerConfig.ReplyTo,
		ReplyToName: providerConfig.ReplyToName,
		TextBody:    textBody,
		HTMLBody:    htmlBody,
	}

	err = sendReq.Execute()
	return
}

type SMSCodeSender struct {
	AppName        string
	Config         config.UserVerifyConfiguration
	SMSClient      sms.Client
	TemplateEngine *template.Engine
}

func (t *SMSCodeSender) Send(verifyCode VerifyCode, userProfile userprofile.UserProfile) (err error) {
	context := prepareVerifyRequestContext(
		verifyCode,
		t.AppName,
		t.Config,
		userProfile,
	)

	var textBody string
	if textBody, err = t.TemplateEngine.ParseTextTemplate(
		authTemplate.VerifyTextTemplateNameForKey(verifyCode.RecordKey),
		context,
		template.ParseOption{Required: true, FallbackTemplateName: authTemplate.TemplateNameVerifySMSText},
	); err != nil {
		return
	}

	err = t.SMSClient.Send(verifyCode.RecordValue, textBody)
	return
}

func prepareVerifyRequestContext(
	verifyCode VerifyCode,
	appName string,
	config config.UserVerifyConfiguration,
	userProfile userprofile.UserProfile,
) map[string]interface{} {
	return map[string]interface{}{
		"appname":       appName,
		"record_key":    verifyCode.RecordKey,
		"record_value":  verifyCode.RecordValue,
		"user_id":       userProfile.ID,
		"user_metadata": userProfile.Data,
		"code":          verifyCode.Code,
		"link": fmt.Sprintf(
			"%s/verify_code_form?code=%s&user_id=%s",
			config.URLPrefix,
			verifyCode.Code,
			userProfile.ID,
		),
	}
}
