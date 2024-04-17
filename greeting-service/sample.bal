import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
};

configurable string name = "Choreo";
configurable string message = "Welcome to Choreo!";
configurable string mountCheck = "Defaultables Mounted";
configurable int testNumber = 10;
configurable int testNumber2 = 20;

// Add a confgiurable to get the configs for a http client
// configurable http:CredentialsConfig clientConfig = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : name, "to" : name, "message" : message, "mountCheck" : mountCheck, "testNumber" : testNumber, "testNumber2" : testNumber2};
        return greetingMessage;
    }
}
