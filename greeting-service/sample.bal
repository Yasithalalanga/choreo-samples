import ballerina/io;
import ballerina/http;

type Tag record {
    string name;
    string description;
};

type Greeting record {
    string name;
    string message;
    Tag[] tags;
};

configurable Greeting[] greetings = [
    {"name" : "Alice", "message" : "Hello Alice!", "tags" : [{"name" : "greeting", "description" : "Greeting message"}]},
    {"name" : "Bob", "message" : "Hello Bob!", "tags" : [{"name" : "greeting", "description" : "Greeting message"}]}
];

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        io:println("Received request for " + name);
        io:print("Greetings in array: ", greetings);
        return greetings[0];
    }
}
