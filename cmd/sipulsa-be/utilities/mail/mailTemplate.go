package mailservice

type MailTemplate struct {
	MailTo  []string
	MailCc  []string
	Subject string
	Body    string
}

func createNewMailTemplate() {

}
