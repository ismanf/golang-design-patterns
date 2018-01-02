package factory

import (
	"errors"
	"fmt"
)

const (
	SMS   = 1
	Email = 2
)

type Notification struct {
	To      string
	Message string
}

type Notifier interface {
	SendNotification(Notification) string
}

type SMSNotification struct{}
func (n SMSNotification) SendNotification(notification Notification) string {
	return fmt.Sprintf("`%s` was sent to number `%s` successfully!", notification.Message, notification.To)
}

type EmailNotification struct{}
func (n EmailNotification) SendNotification(notification Notification) string {
	return fmt.Sprintf("`%s` was sent to email `%s` successfully!", notification.Message, notification.To)
}

// Factory .....
func GetNotifier(notifierType int) (Notifier, error) {
	switch notifierType {
		case SMS:
			return new(SMSNotification), nil
		case Email:
			return new(EmailNotification), nil
		default:
		return nil, errors.New(fmt.Sprintf("Notifier type %d not recognized.", notifierType))
	}
}
