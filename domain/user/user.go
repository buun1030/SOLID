package user

import "fmt"

// Single Responsibility Principle (SRP)

// User represents a user in the system.
type User struct {
	ID       int
	Username string
	Email    string
}

// MongoDBRepository is a concrete implementation of UserRepository
// that interacts with a MongoDB database.
type MongoDBRepository struct{}

func (db *MongoDBRepository) Save(user User) error {
	// Simulate saving user data to a MongoDB database.
	fmt.Printf("Saving user %v to MongoDB\n", user)
	return nil
}

func (db *MongoDBRepository) GetByID(id int) (User, error) {
	// Simulate getting user data from a MongoDB database.
	user := User{
		ID:       id,
		Username: fmt.Sprintf("User from MongoDB with ID %d", id),
		Email:    "mongodb@email.com",
	}
	return user, nil
}

// MySQLRepository is another concrete implementation of UserRepository
// that interacts with a MySQL database.
type MySQLRepository struct{}

func (db *MySQLRepository) Save(user User) error {
	// Simulate saving user data to a MySQL database.
	fmt.Printf("Saving user %v to MySQL\n", user)
	return nil
}

func (db *MySQLRepository) GetByID(id int) (User, error) {
	// Simulate getting user data from a MySQL database.
	user := User{
		ID:       id,
		Username: fmt.Sprintf("User from MySQL with ID %d", id),
		Email:    "mysql@email.com",
	}
	return user, nil
}
