# SOLID
SOLID principles by Go example

## Single Responsibility Principle (SRP): 
This principle states that a class should have only one reason to change. In other words, a class should have only one primary responsibility or job. If a class has multiple responsibilities, it becomes harder to understand and maintain. Breaking down classes into smaller, more focused units makes your code more flexible and easier to modify.

## Open/Closed Principle (OCP): 
The Open/Closed Principle states that software entities (classes, modules, functions, etc.) should be open for extension but closed for modification. This means that you should be able to add new functionality to a system without altering the existing code. You achieve this by using abstractions, interfaces, and polymorphism to allow for future extensions.

## Liskov Substitution Principle (LSP): 
The Liskov Substitution Principle states that objects of a derived class should be able to replace objects of the base class without affecting the correctness of the program. In other words, if you have a base class and a derived class, you should be able to use instances of the derived class wherever you use instances of the base class without causing unexpected behavior.

## Interface Segregation Principle (ISP): 
The Interface Segregation Principle advises that clients should not be forced to depend on interfaces they don't use. In simpler terms, it's better to have multiple small, specific interfaces rather than one large, general-purpose interface. This ensures that classes only need to implement the methods that are relevant to their behavior.

## Dependency Inversion Principle (DIP): 
The Dependency Inversion Principle encourages high-level modules (e.g., classes or components) to depend on abstractions, not on concrete implementations. It also suggests that abstractions should not depend on details; instead, details should depend on abstractions. By following this principle, you can achieve decoupling between components, making it easier to replace or extend parts of your software without affecting the rest of the system.
