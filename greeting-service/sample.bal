import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
    string configurations;
};

configurable string baseUrl = ?;
configurable string containerId = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!" , "configurations" : "baseUrl: " + baseUrl + ", containerId: " + containerId};
        return greetingMessage;
    }
}
