import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';

// Functions bound from Lorca are available on the window object
declare global {
  interface Window {
    add: (a: number, b: number) => Promise<number>;
  }
}
  
window.add(40, 2).then( (result: number) => {
  ReactDOM.render(<App sum={result} />, document.getElementById('root'));
});



// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
