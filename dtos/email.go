package dtos

type Email struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type Message struct {
	From     Email   `json:"from"`
	To       []Email `json:"to"`
	Subject  string  `json:"subject"`
	Text     string  `json:"text"`
	Category string  `json:"category"`
}

type EmailResponse struct {
	Success    bool     `json:"success"`
	MessageIDs []string `json:"message_ids,omitempty"`
	Errors     []string `json:"errors,omitempty"`
}
