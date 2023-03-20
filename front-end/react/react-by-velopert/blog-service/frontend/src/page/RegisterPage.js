import { Helmet } from 'react-helmet-async';

import AuthTemplate from '../component/auth/AuthTemplate';
import RegisterForm from '../container/auth/RegisterForm';
import React from 'react';

const RegisterPage = () => {
  return (
    <AuthTemplate>
      <Helmet>
        <title>Blog Service - Register</title>
      </Helmet>
      <RegisterForm />
    </AuthTemplate>
  );
};

export default RegisterPage;