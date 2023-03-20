import React from 'react';
import { Helmet } from 'react-helmet-async';

import AuthTemplate from '../component/auth/AuthTemplate';
import LoginForm from '../container/auth/LoginForm';

const LoginPage = () => {
  return (
    <AuthTemplate>
      <Helmet>
        <title>Blog Service - Login</title>
      </Helmet>
      <LoginForm />
    </AuthTemplate>
  );
};

export default LoginPage;