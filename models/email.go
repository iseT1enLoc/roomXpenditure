package models

type EmailInformation struct {
	SenderName    string `json:"sender_name"`
	SenderEmail   string `json:"sender_email"`
	ReceiverName  string `json:"receiver_name"`
	ReceiverEmail string `json"receiver_email"`
}
