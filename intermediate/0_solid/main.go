package main

// SRP - Single Responsibility Principle
//    A type, construct, object, should have a single primary responsibility
//    In this example the Journal has the responsibility of storing, adding and removing entries
//    When it comes to persistence we put those somewhere else, so some settings can be shared
// OCP - Open-Closed Principle
//    A type must be open for extension, but closed for modification
//    There should be interfaces that are extendable, but once a type is implemented we shouldn't modify it
//		We should favor expandable interfaces over adding new methods to an existing and already tested type
// LSP - Liskov Substitution Principle
// ISP - Interface Segregation Principle
// DIP - Dependency Inversion Principle

func main() {
	// ...
}
