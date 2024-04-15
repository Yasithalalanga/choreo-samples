import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import oauth from 'axios-oauth-client';

function App() {
  const [response, setResponse] = React.useState(null);

  const fetchData = async () => {
    const serviceURL = window?.configs?.serviceURL || "/";
    const consumerKey = window?.configs?.consumerKey || "";
    const consumerSecret = window?.configs?.consumerSecret || "";
    const tokenURL = window?.configs?.tokenURL || "";

    console.log('serviceURL', serviceURL);
    console.log('consumerKey', consumerKey);
    console.log('consumerSecret', consumerSecret);
    console.log('tokenURL', tokenURL);

    const getClientCredentials = oauth.clientCredentials(
      axios.create(),
      tokenURL,
      consumerKey,
      consumerSecret
    );

    try {
      const auth = await getClientCredentials();
      const accessToken = auth.access_token;
      console.log('accessToken', accessToken);

      let resourcePath = '?name=hello';
      const apiResponse = await axios.get(serviceURL + resourcePath, {
        headers: {
          Authorization: `Bearer ${accessToken}`
        }
      });

      setResponse(apiResponse.data);
      console.log('response', apiResponse.data);
    } catch (error) {
      console.error('API call failed:', error);
      setResponse({ error: 'API call failed. Check the console for more information.' });
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <button onClick={fetchData} className="App-link">
          Fetch Data
        </button>
        {response && <pre>{JSON.stringify(response, null, 2)}</pre>}
        <a
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
          className="App-link"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
