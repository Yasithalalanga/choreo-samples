import ballerina/http;
import ballerina/io;


configurable string[] requiredNames = ?;
configurable string[] optionalNames = [];
configurable Greeting[] greetings = [];

type Greeting record {
    string 'from;
    string to;
    string message;
};

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        io:println("Received request from: " + name);
        io:println("Required names: " + requiredNames.toBalString());
        io:println("Optional names: " + optionalNames.toBalString());
        io:println("Greetings: " + greetings.toBalString());
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        return greetingMessage;
    }
}
