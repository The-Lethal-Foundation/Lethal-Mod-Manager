import React from 'react';
import logo from './logo.svg';
import './App.css';

declare global {
  interface Window {
    add: (a: number, b: number) => Promise<number>
    process: () => Promise<any>
  }
}

window.process().then((obj) => {
  console.log(obj)
})

function myAdd(a: number, b: number):number {
  return a + b
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
