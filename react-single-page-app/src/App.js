import logo from './logo.svg';
import './App.css';

import oauth from 'axios-oauth-client'
import axios from 'axios';

async function App() {
  const serviceURL = window?.configs?.serviceURL ? window.configs.serviceURL : "/";
  const consumerKey = window?.configs?.consumerKey ? window.configs.consumerKey : "";
  const consumerSecret = window?.configs?.consumerSecret ? window.configs.consumerSecret : "";
  const tokenURL = window?.configs?.tokenURL ? window.configs.tokenURL : "";

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
  const auth = await getClientCredentials();
  const accessToken = auth.access_token;

  console.log('accessToken', accessToken);

  let resourcePath = '?name=hello';
  const response = await axios.get(serviceURL + resourcePath, {
    headers: {
      Authorization: `Bearer ${accessToken}`
    }
  });

  console.log('response', response);


  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
