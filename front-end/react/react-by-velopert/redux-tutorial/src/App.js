import React from 'react';
import CounterContainer from './container/CounterContainer';
import TodosContainer from './container/TodosContainer';

function App() {
  return (
    <div style={{padding: '10px'}}>
      <CounterContainer />
      <hr/>
      <TodosContainer/>
    </div>
  );
}

export default App;
