import ballerina/http;

configurable boolean isAdmin = ?;
configurable byte age = ?;
configurable int port = ?;

configurable float height = ?;
configurable decimal salary = ?;

configurable string name = ?;

enum Country {
  LK = "Sri Lanka",
  US = "United States"
}
configurable Country country = ?;

configurable float|int|string measurement = ?;

service / on new http:Listener(8090) {
    resource function get greeting() returns string {
        string greetingMessage = "Hello, World!";

        if (isAdmin) {
            greetingMessage = greetingMessage + " Is Admin Set to True";
        } else {
            greetingMessage = greetingMessage + " Is Admin Set to False";
        }

        // Append the configurable values with keys
        greetingMessage = greetingMessage + " Age: " + age.toBalString();
        greetingMessage = greetingMessage + " Port: " + port.toBalString();
        greetingMessage = greetingMessage + " Height: " + height.toBalString();
        greetingMessage = greetingMessage + " Salary: " + salary.toBalString();

        greetingMessage = greetingMessage + " Name: " + name;

        greetingMessage = greetingMessage + " Country: " + country.toString();
        greetingMessage = greetingMessage + " Measurement: " + measurement.toBalString();

        return greetingMessage;
    }
}
