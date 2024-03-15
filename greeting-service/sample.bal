import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
};

configurable string name = "Choreo";
configurable string message = "Welcome to Choreo!";
configurable string mountCheck = ?;
configurable string sampleConfig = ?;

// Add a confgiurable to get the configs for a http client
// configurable http:CredentialsConfig clientConfig = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!", "mountCheck" : mountCheck, "sampleConfig" : sampleConfig};
        return greetingMessage;
    }
}
