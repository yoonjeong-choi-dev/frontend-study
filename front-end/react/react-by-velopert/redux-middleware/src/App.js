import React from 'react';
import CounterContainer from './container/counter/CounterContainer';
import CounterContainerWithThunk from './container/counter/CounterContainerWithThunk';
import JsonHolderContainer from './container/JsonHolderContainer';
import CounterContainerWithSaga from './container/counter/CounterContainerWithSaga';
import JsonHolderContainerForSaga from './container/JsonHolderContainerForSaga';

function App() {
  return (
    <div style={{padding: '10px'}}>
      <CounterContainer/>
      <CounterContainerWithThunk/>
      <CounterContainerWithSaga/>
      <hr/>
      <div style={{display: 'grid', gap: '1rem', gridTemplateColumns: '1fr 1fr'}}>
        <JsonHolderContainer/>
        <JsonHolderContainerForSaga/>
      </div>

    </div>
  );
}

export default App;
