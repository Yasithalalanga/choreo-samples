import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
};

configurable string username = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!", "username": username};
        return greetingMessage;
    }
}
