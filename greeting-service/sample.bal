import ballerina/http;

configurable string name = ?;
configurable string message = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns string {
        string greetingMessage = "Hello, " + name + "!";
        greetingMessage = greetingMessage + "  message: " + message;
        return greetingMessage;
    }
}
