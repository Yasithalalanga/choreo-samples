import ballerina/http;
import ballerina/io;

type Greeting record {
    string 'from;
    string to;
    string message;
};

type ClientsiteConfig record {
    string clientUrl;
    string certPath;
};

// Get configurations from config file or environment variables
configurable ClientsiteConfig clientsiteConfig = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        // Just print the values from mobileapp_api
        io:println("Client URL: " + clientsiteConfig.clientUrl);
        io:println("Cert Path: " + clientsiteConfig.certPath);
        return greetingMessage;
    }
}
