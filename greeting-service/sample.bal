import ballerina/http;
import ballerina/io;

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
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        io:println("Retry Config: ", testRetryConfig);
        return greetingMessage;
    }
}
