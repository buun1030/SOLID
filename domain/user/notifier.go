package user

import (
	"fmt"
)

// EmailNotifier is an implementation of the Notifier interface for sending emails.
type EmailNotifier struct {
	EmailRecipient string
}

func (en EmailNotifier) SendNotification(message string) error {
	// Simulate sending an email.
	fmt.Printf("Sending email to %s: %s\n", en.EmailRecipient, message)
	return nil
}

// SMSNotifier is another implementation of the Notifier interface for sending SMS.
type SMSNotifier struct {
	PhoneNumber string
}

func (sn SMSNotifier) SendNotification(message string) error {
	// Simulate sending an SMS.
	fmt.Printf("Sending SMS to %s: %s\n", sn.PhoneNumber, message)
	return nil
}
