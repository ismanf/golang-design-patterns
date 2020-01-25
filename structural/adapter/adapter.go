package adapter

//old, bad service which we can't change
type OldEmailSender interface {
	Send(from, to, subject, body string) error
}

type OldEmailSvc struct{}

func (s OldEmailSvc) Send(from, to, subject, body string) error {
	//Send email...
	return nil
}

//new service which has state
//we have a client only supporting this interface
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

//adapter for the old email service
type OldEmailServiceAdapter struct {
	From, To, Subject, Body string
	OldEmailService         OldEmailSender
}

func (a OldEmailServiceAdapter) Send() error {
	return a.OldEmailService.Send(a.From, a.To, a.Subject, a.Body)
}

/*
    client which only supports the new email service interface
	EmailClient{ Mail: OldEmailServiceAdapter{...}} old service
	EmailClient{ Mail: NewEmailService{...}} new service
*/
type EmailClient struct {
	Mail NewEmailSender
}

func (e EmailClient) SendEmail(From, To, Subject, Body string) error {
	return e.Mail.Send()
}
