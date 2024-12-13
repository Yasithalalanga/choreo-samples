import ballerina/http;
import ballerina/io;


configurable string[] requiredNames = ?;
configurable string[] optionalNames = ["Choreo", "Ballerina"];
configurable Greeting[] greetings = [];

// Add a required greetings array configuration
configurable Greeting[] requiredGreetings = ?;
// Add a required nested array configuration
configurable Greeting[][] nestedGreetings = ?;
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
        io:println("Required greetings: " + requiredGreetings.toBalString());

        // Add a value to the nestedGreetings array

        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        return greetingMessage;
    }
}
