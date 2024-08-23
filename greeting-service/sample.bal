import ballerina/io;

configurable string name = "World";
configurable string brand = "Choreo";
configurable string message = "Welcome to Choreo!";

public function main() {
    io:println("Hello, World!");
    io:println("From: " + brand);
    io:println("To: " + name);
    io:println("Message: " + message);
}
