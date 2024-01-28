import React, { useState } from 'react';
import './App.css';

const App: React.FC = () => {

  const [active, setActive] = useState<boolean>(true);

  return (
    <div className="App">
      <header className="App-header">
        
        <div className='text-red-500'>
          ABC 123
        </div>

      </header>
    </div>
  );
}

export default App;
