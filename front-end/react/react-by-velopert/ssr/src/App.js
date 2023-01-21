import React from 'react';
import Menu from './component/Menu';
import { Route, Routes } from 'react-router-dom';
import RedPage from './page/RedPage';
import BluePage from './page/BluePage';

function App() {
  return (
    <div>
      <h1>Server Side Rendering</h1>
      <Menu/>
      <Routes>
        <Route path={'/red'} element={<RedPage/>}/>
        <Route path={'/blue'} element={<BluePage/>}/>
      </Routes>
    </div>
  );
}

export default App;
