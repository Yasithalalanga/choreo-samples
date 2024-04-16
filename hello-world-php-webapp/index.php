<?php
// Environment variables - these should be set in your environment or configuration file
$serviceUrl = getenv('SVC_URL');
$consumerKey = getenv('CONSUMER_KEY');
$consumerSecret = getenv('CONSUMER_SECRET');
$tokenUrl = getenv('TOKEN_URL');

// Display the environment variable values
echo '<h2>Environment Variables</h2>';
echo '<p>Service URL: ' . htmlspecialchars($serviceUrl) . '</p>';
echo '<p>Consumer Key: ' . htmlspecialchars($consumerKey) . '</p>';
echo '<p>Consumer Secret: ' . htmlspecialchars($consumerSecret) . '</p>';
echo '<p>Token URL: ' . htmlspecialchars($tokenUrl) . '</p>';

// Function to generate a token
function generateToken($tokenUrl, $consumerKey, $consumerSecret) {
    $curl = curl_init();
    curl_setopt($curl, CURLOPT_URL, $tokenUrl);
    curl_setopt($curl, CURLOPT_POST, true);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($curl, CURLOPT_HTTPHEADER, [
        'Content-Type: application/x-www-form-urlencoded',
    ]);
    curl_setopt($curl, CURLOPT_POSTFIELDS, http_build_query([
        'grant_type' => 'client_credentials',
        'client_id' => $consumerKey,
        'client_secret' => $consumerSecret,
    ]));

    $response = curl_exec($curl);
    if (!$response) {
        die('Error: "' . curl_error($curl) . '" - Code: ' . curl_errno($curl));
    }
    curl_close($curl);
    
    $response = json_decode($response, true);
    return $response['access_token'] ?? '';
}

$token = generateToken($tokenUrl, $consumerKey, $consumerSecret);

// Function to call the service URL with the token
function callService($serviceUrl, $token) {
    $query = '?name=hello'; // Appending the query string
    $curl = curl_init();
    curl_setopt($curl, CURLOPT_URL, $serviceUrl . $query);
    curl_setopt($curl, CURLOPT_HTTPHEADER, [
        'Authorization: Bearer ' . $token,
    ]);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

    $response = curl_exec($curl);
    if (!$response) {
        die('Error: "' . curl_error($curl) . '" - Code: ' . curl_errno($curl));
    }
    curl_close($curl);

    return $response;
}

$response = callService($serviceUrl, $token);

echo '<h2>Response from Service</h2>';
echo '<pre>' . htmlspecialchars($response) . '</pre>';
?>
