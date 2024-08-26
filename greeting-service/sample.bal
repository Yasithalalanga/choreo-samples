// import ballerina/http;

// configurable string nameV2 = "Choreo";
// configurable string nameV1 = "Choreo";
// configurable string name-under="vjrpods";

// service / on new http:Listener(8090) {
//     resource function get .(string name) returns string {
//         string greetingMessage = "Hello, " + name + "!";
//         greetingMessage = greetingMessage + " Welcome to " + nameV1 + "!";
//         greetingMessage = greetingMessage + " Welcome to " + nameV2 + "!";
//         return greetingMessage;
//     }
// }

import ballerina/io;

public function main() {
    io:println("Hello, World!");
}
