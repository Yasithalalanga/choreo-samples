import ballerina/http;

configurable string name = ?;
configurable string message = ?;
configurable Greeting greeting = ?;
configurable string[] hobbies = ?;
configurable string sampleDefault = "default";

type Greeting record {|
    string name;
    string message;
|};

service / on new http:Listener(8090) {
    resource function get .(string name) returns string {
        string greetingMessage = "Hello, " + name + "!";
        greetingMessage = greetingMessage + "  message: " + message;

        greetingMessage = greetingMessage + "  name: " + greeting.name;
        greetingMessage = greetingMessage + "  message: " + greeting.message;

        greetingMessage = greetingMessage + "  hobbies: " + hobbies[0];

        greetingMessage = greetingMessage + "  sampleDefault: " + sampleDefault;

        return greetingMessage;
    }
}

