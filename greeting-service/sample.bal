import ballerina/http;
import ballerina/os;
import ballerina/io;

type Greeting record {
    string 'from;
    string to;
    string message;
};

service / on new http:Listener(8090) {
    resource function get .(string name) returns Greeting {
        Greeting greetingMessage = {"from" : "Choreo", "to" : name, "message" : "Welcome to Choreo!"};

        // Read environment variables

        string serviceUrl = os:getEnv("SVC_URL");
        string consumerKey = os:getEnv("CONSUMER_KEY");
        string consumerSecret = os:getEnv("CONSUMER_SECRET");
        string tokenUrl = os:getEnv("TOKEN_URL");

        // Print environment variables
        io:println("Environment variables:");
        io:println("Service URL: " + serviceUrl);
        io:println("Consumer Key: " + consumerKey);
        io:println("Consumer Secret: " + consumerSecret);
        io:println("Token URL: " + tokenUrl);


          http:Client httpClient = check new http:Client(serviceURL, {
            auth: {
                tokenUrl: tokenURL,
                clientId: consumerKey,
                clientSecret: consumerSecret
            }
        });

        var response = check httpClient->get("/");
        io:println("Response: " + response);

        return greetingMessage;
    }
}
