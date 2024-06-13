import ballerina/http;
import ballerina/io;

type Greeting record {
    string 'from;
    string to;
    string message;
};

type Calc record {|
    decimal amount;
    Dates day;
|};

// Dates Enum
public enum Dates {
    SUNDAY = "Sunday",
    MONDAY = "Monday",
    TUESDAY = "Tuesday",
    WEDNESDAY = "Wednesday",
    THURSDAY = "Thursday",
    FRIDAY = "Friday",
    SATURDAY = "Saturday"
}

configurable decimal amount = ?;
configurable Calc Calculation = ?;
configurable Dates day = ?;

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};
        io:println("Decimal amount: " + amount.toBalString());
        io:println(amount);

        io:println("Object Calc: " + Calculation.amount.toBalString());
        io:println(Calculation.amount);
        io:println("Object Calc: " + Calculation.day.toString());
        io:println(Calculation.day);

        io:println("Enum Day: " + day.toString());
        io:println(day);
        return greetingMessage;
    }
}
