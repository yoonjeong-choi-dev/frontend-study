import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';
import { HelmetProvider } from 'react-helmet-async';
import { applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import { composeWithDevTools } from 'redux-devtools-extension';
import './index.css';

import App from './App';
import rootReducer from './redux';
import { check, tempSetUser } from './redux/auth/user';

const store = createStore(rootReducer,
  composeWithDevTools(applyMiddleware(thunk)),
);

function loadUser() {
  try {
    const user = localStorage.getItem('user');
    if (!user) return;

    store.dispatch(tempSetUser(JSON.parse(user)));
    store.dispatch(check());
  } catch (e) {
    console.error('local storage is not working');
  }
}

loadUser();
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Provider store={ store }>
      <HelmetProvider>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </HelmetProvider>
    </Provider>
  </React.StrictMode>,
);

