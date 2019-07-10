import React from 'react';
import logo from './logo.svg';
import './App.css';

// Functions bound from Lorca are available on the window object
declare global {
  interface Window {
    add: (a: number, b: number) => Promise<number>;
  }
}

// Hack check to make silly App.test.tsx work
if(window.add) {
  window.add(40, 2).then( (result: number) => {
    console.log(result);
  });
}

const App: React.FC = () => {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload!
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
