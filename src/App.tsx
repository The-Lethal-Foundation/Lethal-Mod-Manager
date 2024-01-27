import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';

interface AppProps {
  sum: number
}

const App: React.FC<AppProps> = (props: AppProps) => {

  const [active, setActive] = useState<boolean>(true);

  return (
    <div className="App">
      <header className="App-header">
        
        {active && 
          <img src={logo} className="App-logo" alt="logo" />
        }
        <p>
          Edit <code>src/App.tsx</code> and save to reload!
        </p>
        <p>
          Sum from Go: {props.sum}
        </p>
        <a
          className="App-link"
          href="https://www.linkedin.com/in/aniervs/"
          target="_blank"
          rel="noopener noreferrer"
        >
          What's up Anier!
        </a>

        <button  onClick={() => setActive(!active)}>Toggle</button>
      </header>
    </div>
  );
}

export default App;
