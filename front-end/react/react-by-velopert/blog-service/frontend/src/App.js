import React from 'react';
import { Route, Routes } from 'react-router-dom';

import RegisterPage from './page/RegisterPage';
import LoginPage from './page/LoginPage';
import PostPage from './page/PostPage';
import PostListPage from './page/PostListPage';
import WritePage from './page/WritePage';

function App() {
  return (
    <Routes>
      <Route path='/' element={ <PostListPage /> } />
      <Route path='/login' element={ <LoginPage /> } />
      <Route path='/register' element={ <RegisterPage /> } />
      <Route path='/write' element={ <WritePage /> } />
      <Route path='/:username'>
        <Route index element={ <PostListPage /> } />
        <Route path=':postId' element={ <PostPage /> } />
      </Route>
    </Routes>
  );
}

export default App;
