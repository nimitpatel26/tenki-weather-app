/**
 * This file contains the entry point for
 * the React application
 */
import React from 'react';
import ReactDOM from 'react-dom';
import App from './app/app';
import './index.sass';


ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);
