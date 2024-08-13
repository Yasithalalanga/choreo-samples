import ballerina/http;

// configurable boolean isAdmin = ?;
// configurable byte age = ?;
// configurable int port = ?;

// configurable float height = ?;
// configurable decimal salary = ?;

// configurable string name = ?;

// enum Country {
//   LK = "Sri Lanka",
//   US = "United States"
// }
// configurable Country country = ?;
// configurable float|int|string measurement = ?;

// type Greeter record {|
//     string name;
//     string message;
// |};

// Defaultable fields
configurable string defaultName = "John Doe";
configurable string defaultMessage = "Hello, World!";

// configurable Greeter greeter = ?;
// configurable string[] hobbies = ?;
// configurable int[] grades = ?;

// Nested Objects

// type Event record {|
//     Greeter greeter;
//     string eventName;
// |};

// configurable Event event = ?;

service / on new http:Listener(8090) {
    resource function get greeting() returns string {
        string greetingMessage = "Hello, World!";

        // if (isAdmin) {
        //     greetingMessage = greetingMessage + " Is Admin Set to True";
        // } else {
        //     greetingMessage = greetingMessage + " Is Admin Set to False";
        // }

        // // Append the configurable values with keys
        // greetingMessage = greetingMessage + " Age: " + age.toBalString();
        // greetingMessage = greetingMessage + " Port: " + port.toBalString();
        // greetingMessage = greetingMessage + " Height: " + height.toBalString();
        // greetingMessage = greetingMessage + " Salary: " + salary.toBalString();

        // greetingMessage = greetingMessage + " Name: " + name;

        // greetingMessage = greetingMessage + " Country: " + country.toString();
        // greetingMessage = greetingMessage + " Measurement: " + measurement.toBalString();

        // greetingMessage = greetingMessage + "Greeter Name: " + greeter.name;
        // greetingMessage = greetingMessage + "Greeter Message: " + greeter.message;

        // Append the defaultable values with keys
        greetingMessage = greetingMessage + " Default Name: " + defaultName;
        greetingMessage = greetingMessage + " Default Message: " + defaultMessage;

        // // Append the array values with keys
        // greetingMessage = greetingMessage + " Hobbies: " + hobbies.toString();
        // greetingMessage = greetingMessage + " Grades: " + grades.toString();

        // // Append the nested object values with keys
        // greetingMessage = greetingMessage + " Event Name: " + event.eventName;
        // greetingMessage = greetingMessage + " Event Greeter Name: " + event.greeter.name;
        // greetingMessage = greetingMessage + " Event Greeter Message: " + event.greeter.message;

        return greetingMessage;
    }
}
