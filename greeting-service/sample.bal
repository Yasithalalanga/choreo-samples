import ballerina/http;

type Greeting record {
    string 'from;
    string to;
    string message;
};

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        return greetingMessage;
    }

    // Add a new get resource whihc prints lifecycle change successful message on  /lifecycle path
    resource function get lifecycle() returns string {
        return "Lifecycle change successful!";
    }
}
