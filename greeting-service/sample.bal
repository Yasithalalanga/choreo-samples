// import ballerina/http;

// type Greeting record {
//     string 'from;
//     string to;
//     string message;
// };

// service / on new http:Listener(8090) {
//     resource function get .(string name) returns Greeting {
//         Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
//         return greetingMessage;
//     }
// }

import ballerina/io;

configurable string name = ?;
configurable string brand = ?;
configurable string message = ?;

public function main() {
    io:println("Hello, World!");
    io:println("From: " + brand);
    io:println("To: " + name);
    io:println("Message: " + message);
}
