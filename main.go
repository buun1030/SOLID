package main

import (
	"fmt"
	"solid/domain/document"
	"solid/domain/user"
)

func main() {
	// SRP Example
	userRepository := &user.MongoDBRepository{}
	userService := user.UserRegistrationService{Repo: userRepository}
	err := userService.Register("john_doe", "john@example.com")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// OCP Example
	emailNotifier := user.EmailNotifier{EmailRecipient: "admin@example.com"}
	smsNotifier := user.SMSNotifier{PhoneNumber: "+1234567890"}

	// LSP Example
	notifyingUserServiceEmail := user.NotifyingUserService{
		UserRegistrationService: userService,
		Notifier:                emailNotifier,
	}
	notifyingUserServiceSMS := user.NotifyingUserService{
		UserRegistrationService: userService,
		Notifier:                smsNotifier,
	}

	notifyingUserServiceEmail.RegisterAndNotify("alice_doe", "alice@example.com", "Welcome, Alice!")
	notifyingUserServiceSMS.RegisterAndNotify("bob_doe", "bob@example.com", "Welcome, Bob!")

	reader := document.TextDocument{}

	// Client interacts with specific interfaces, avoiding unnecessary methods.
	document.ClientViewEvidence(reader)
}
