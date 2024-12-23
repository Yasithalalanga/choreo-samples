import ballerina/http;
import ballerina/io;

type Greeting record {
    string 'from;
    string to;
    string message;
};

type ClientSiteConfig record {
    string clientUrl;
    string certPath;
};

// Get configurations from config file or environment variables
configurable ClientSiteConfig clientSiteConfig = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        // Just print the values from mobileapp_api
        io:println("Client URL: " + clientSiteConfig.clientUrl);
        io:println("Cert Path: " + clientSiteConfig.certPath);
        return greetingMessage;
    }
}
