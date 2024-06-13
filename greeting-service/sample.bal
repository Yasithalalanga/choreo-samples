import ballerina/http;
import ballerina/io;

type Greeting record {
    string 'from;
    string to;
    string message;
};

type Calc record {|
    decimal amount;
|};

configurable decimal amount = ?;
configurable Calc Calculation = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        io:println("Decimal amount: " + amount.toBalString());
        io:println(amount);

        io:println("Object Calc: " + Calculation.amount.toBalString());
        io:println(Calculation.amount);
        return greetingMessage;
    }
}
