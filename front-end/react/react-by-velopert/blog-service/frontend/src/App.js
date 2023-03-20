import React from 'react';
import { Route, Routes } from 'react-router-dom';
import {Helmet} from 'react-helmet-async';

import RegisterPage from './page/RegisterPage';
import LoginPage from './page/LoginPage';
import PostPage from './page/PostPage';
import PostListPage from './page/PostListPage';
import WritePage from './page/WritePage';
import EditPage from './page/EditPage';

function App() {
  return (
    <>
      <Helmet>
        <title>Blog Service</title>
      </Helmet>
      <Routes>
        <Route path='/' element={ <PostListPage /> } />
        <Route path='/login' element={ <LoginPage /> } />
        <Route path='/register' element={ <RegisterPage /> } />
        <Route path='/write' element={ <WritePage /> } />
        <Route path='/edit' element={ <EditPage /> } />
        <Route path='/:username'>
          <Route index element={ <PostListPage /> } />
          <Route path=':postId' element={ <PostPage /> } />
        </Route>
      </Routes>
    </>
  );
}

export default App;
