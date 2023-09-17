package user

// Open/Closed Principle (OCP)

// Notifier interface and its sending notifications implementations
// (EmailNotifier and SMSNotifier)
// allow for easy extension by adding new notification methods
// without modifying existing code.
type Notifier interface {
	SendNotification(message string) error
}

// Liskov Substitution Principle (LSP)

// derived classes (EmailNotifier and SMSNotifier)
// can replace objects of the base class (Notifier)
// without affecting the correctness of the program
type NotifyingUserService struct {
	UserRegistrationService
	Notifier
}

func (nus NotifyingUserService) RegisterAndNotify(username, email, message string) error {
	err := nus.UserRegistrationService.Register(username, email)
	if err != nil {
		return err
	}
	return nus.SendNotification(message)
}

// UserRepository is a low-level module that interacts with user data.
type UserRepository interface {
	Save(user User) error
	GetByID(id int) (User, error)
}

// Single Responsibility Principle (SRP)

// Could be bad practice to have more than one responsibility.

// UserService is a high-level module that depends on UserRepository through the abstraction.
type UserService struct {
	Repo UserRepository
}

func (us UserService) Register(username, email string) error {
	user := User{Username: username, Email: email}
	return us.Repo.Save(user)
}

func (us UserService) GetUser(id int) (User, error) {
	user, err := us.Repo.GetByID(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (us UserService) UpdateUserProfile(id int, newUsername, newEmail string) error {
	return nil
}

// Could be better practice to have one responsibility per type.

type UserRegistrationService struct {
	Repo UserRepository
	// more maintainable to add more dependencies as necessary
}

func (us UserRegistrationService) Register(username, email string) error {
	user := User{Username: username, Email: email}
	return us.Repo.Save(user)
}

type UserRetrievalService struct {
	Repo UserRepository
	// more maintainable to add more dependencies as necessary
}

func (us UserRetrievalService) GetUser(id int) (User, error) {
	user, err := us.Repo.GetByID(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

type UserServiceToInterface interface {
	Register(username, email string) error
	GetUser(id int) (User, error)
}
