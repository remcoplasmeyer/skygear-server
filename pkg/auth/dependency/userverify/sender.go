package userverify

import (
	"errors"
	"fmt"

	"github.com/go-gomail/gomail"
	"github.com/skygeario/skygear-server/pkg/auth/dependency/userprofile"
	"github.com/skygeario/skygear-server/pkg/core/config"
	"github.com/skygeario/skygear-server/pkg/core/mail"
	"github.com/skygeario/skygear-server/pkg/core/template"
)

type CodeSender interface {
	Send(userProfile userprofile.UserProfile) error
}

type EmailCodeSender struct {
	Key     string
	AppName string
	Config  config.UserVerifyConfiguration
	Dialer  *gomail.Dialer
	CodeGenerator
}

func (e *EmailCodeSender) Send(userProfile userprofile.UserProfile) (err error) {
	var recordValue string
	var ok bool
	if recordValue, ok = userProfile.Data[e.Key].(string); !ok {
		return errors.New(e.Key + " is invalid in user data")
	}

	var keyConfig config.UserVerifyKeyConfiguration
	if keyConfig, ok = e.Config.ConfigForKey(e.Key); !ok {
		return errors.New("provider for " + e.Key + " not found")
	}

	providerConfig := keyConfig.ProviderConfig

	code := e.CodeGenerator.Generate()
	context := map[string]interface{}{
		"appname":      e.AppName,
		"record_key":   e.Key,
		"record_value": recordValue,
		"user_id":      userProfile.RecordID,
		"user":         userProfile.ToMap(),
		"code":         code,
		"link": fmt.Sprintf(
			"%s/auth/verify-code/form?code=%s&user_id=%s",
			e.Config.URLPrefix,
			code,
			userProfile.RecordID,
		),
	}

	var textBody string
	if textBody, err = template.ParseTextTemplateFromURL(providerConfig.TextURL, context); err != nil {
		return
	}

	var htmlBody string
	if providerConfig.HTMLURL != "" {
		if htmlBody, err = template.ParseHTMLTemplateFromURL(providerConfig.HTMLURL, context); err != nil {
			return
		}
	}

	sendReq := mail.SendRequest{
		Dialer:      e.Dialer,
		Sender:      providerConfig.Sender,
		SenderName:  providerConfig.SenderName,
		Recipient:   recordValue,
		Subject:     providerConfig.Subject,
		ReplyTo:     providerConfig.ReplyTo,
		ReplyToName: providerConfig.ReplyToName,
		TextBody:    textBody,
		HTMLBody:    htmlBody,
	}

	err = sendReq.Execute()
	return
}
