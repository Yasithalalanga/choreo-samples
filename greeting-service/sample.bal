import ballerina/http;
import ballerina/io;

type Greeting record {
    string 'from;
    string to;
    string message;
};

configurable  string[] username = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!", "username": username.toString()};
        // Iterate and print the username
        foreach var user in username {
            io:println("Username: " + user);
        }
        return greetingMessage;
    }
}
