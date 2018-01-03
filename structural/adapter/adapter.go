package adapter

//old bad service which we can't chane
type OldEmailSender interface {
	Send(from, to, subject, body string) error
}

type OldEmailSvc struct{}

func (s OldEmailSvc) Send(from, to, subject, body string) error {
	//Send email...
	return nil
}

//new service ehich has state
//we have client only supports this interface
type NewEmailSender interface {
	Send() error
}

type NewEmailService struct {
	From, To, Subject, Body string
}

func (s NewEmailService) Send() error {
	//state already initialized just send
	return nil
}

//adapter for old email service
type OldEmailServiceAdapter struct {
	From, To, Subject, Body string
	OldEmailService         OldEmailSender
}

func (a OldEmailServiceAdapter) Send() error {
	return a.OldEmailService.Send(a.From, a.To, a.Subject, a.Body)
}

/*
    client which only supports new email service interface
	EmailClient{ Mail: OldEmailServiceAdapter{...}} old service
	EmailClient{ Mail: NewEmailService{...}} new service
*/
type EmailClient struct {
	Mail NewEmailSender
}

func (e EmailClient) SendEmail(From, To, Subject, Body string) error {
	return e.Mail.Send()
}
