import ballerina/http;
import ballerina/io;

type Greeting record {
    string 'from;
    string to;
    string message;
};

public configurable readonly & map<string> companyEmailRecipientsMapping = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        io:println("Listing the company-email-recepient-list: ");
        io:println(companyEmailRecipientsMapping);
        return greetingMessage;
    }
}
