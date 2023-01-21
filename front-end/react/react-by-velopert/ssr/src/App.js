import React from 'react';
import Menu from './component/Menu';
import { Route, Routes } from 'react-router-dom';
import RedPage from './page/RedPage';
import BluePage from './page/BluePage';
import UsersPage from './page/UsersPage';

function App() {
  return (
    <div>
      <h1>Server Side Rendering</h1>
      <Menu/>
      <Routes>
        <Route path={'/red'} element={<RedPage/>}/>
        <Route path={'/blue'} element={<BluePage/>}/>
        <Route path={'/users/*'} element={<UsersPage/>}/>
      </Routes>
    </div>
  );
}

export default App;
