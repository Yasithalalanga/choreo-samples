import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
};

configurable string name = "Choreo";
configurable string message = "Welcome to Choreo!";

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        return greetingMessage;
    }
}
