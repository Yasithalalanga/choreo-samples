import ballerina/http;

configurable string name = ?;
configurable string message = ?;
configurable Greeting gr = ?;
configurable string[] hobbies = ?;

type Greeting record {|
    string name;
    string message;
|};

service / on new http:Listener(8090) {
    resource function get .(string name) returns string {
        string greetingMessage = "Hello, " + name + "!";
        greetingMessage = greetingMessage + "  message: " + message;

        greetingMessage = greetingMessage + "  name: " + gr.name;
        greetingMessage = greetingMessage + "  message: " + gr.message;

        greetingMessage = greetingMessage + "  hobbies: " + hobbies[0];

        return greetingMessage;
    }
}

