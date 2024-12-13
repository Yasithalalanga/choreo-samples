import ballerina/http;
import ballerina/io;


configurable string[] requiredNames = ?;
configurable string[] optionalNames = ["Choreo", "Ballerina"];
configurable Greeting[] greetings = [];

type RetryConfig record {|
    *http:RetryConfig;
|};

configurable RetryConfig testRetryConfig = ?;

http:Client testClient = check new ("testUrl", {
    retryConfig: {
        ...testRetryConfig
    }
});

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
