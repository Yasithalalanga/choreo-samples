import ballerina/io;

// Required fields
configurable boolean isAdmin = ?;
configurable byte age = ?;
configurable int port = ?;

// Defaultable fields
configurable string defaultName = "John Doe";
configurable string defaultMessage = "Hello, World!";

public function main() {
    io:println("Hello, World!");
    io:println("Is Admin: " + isAdmin.toBalString());
    io:println("Age: " + age.toBalString());
    io:println("Port: " + port.toBalString());
    io:println("Default Name: " + defaultName);
    io:println("Default Message: " + defaultMessage);
}
