import React from 'react';
import Menu from './component/box/Menu';
import { Route, Routes } from 'react-router-dom';
import loadable from '@loadable/component';

// import RedPage from './page/box/RedPage';
// import BluePage from './page/box/BluePage';
// import UsersPage from './page/UsersPage';
const RedPage = loadable(() => import('./page/box/RedPage'));
const BluePage = loadable(() => import('./page/box/BluePage'));
const UsersPage = loadable(() => import('./page/UsersPage'));

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
