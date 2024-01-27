import React, { useState } from 'react';
import './App.css';

const App: React.FC = () => {

  const [active, setActive] = useState<boolean>(true);

  return (
    <div className="App">
      <header className="App-header">
        
        I love Angelina {active ? '❤️' : '💔'}

        <button onClick={() => setActive(!active)}>
          {active ? 'Break up' : 'Get back together'}
        </button>

      </header>
    </div>
  );
}

export default App;
