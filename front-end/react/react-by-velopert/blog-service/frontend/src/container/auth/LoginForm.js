import { useDispatch, useSelector } from 'react-redux';
import { useEffect, useState } from 'react';

import { changeField, initializeForm, login } from '../../redux/auth/auth';
import AuthForm from '../../component/auth/AuthForm';
import { useNavigate } from 'react-router-dom';
import { check } from '../../redux/auth/user';

const formName = 'login';

const LoginForm = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const { form, auth, authError, user } = useSelector((state) => ({
    form: state.auth[formName],
    auth: state.auth.auth,
    authError: state.auth.authError,
    user: state.user.user,
  }));

  const [error, setError] = useState(null);

  const onChange = (e) => {
    const { name, value } = e.target;
    dispatch(
      changeField({
        form: formName,
        key: name,
        value,
      }),
    );
  };

  const onSubmit = (e) => {
    e.preventDefault();
    const { name, password } = form;
    if ([name, password].includes('')) {
      setError('모든 정보를 입력해주세요');
      return;
    }
    dispatch(login({ name, password }));
  };

  useEffect(() => {
    dispatch(initializeForm(formName));
  }, [dispatch]);

  useEffect(() => {
    if (authError) {
      console.log('Error for login: ', authError);
      const code = authError.response.status;
      if (code === 401) {
        setError('일치하는 회원 정보가 없습니다');
      } else {
        setError('로그인 실패');
      }
      return;
    }

    if (auth) {
      console.log('Success to login: ', auth);
      dispatch(check());
    }
  }, [auth, authError, dispatch]);

  useEffect(() => {
    if (user) {
      console.log('Success to call /auth/check', user);
      navigate('/');
      try {
        localStorage.setItem('user', JSON.stringify(user));
      } catch (e) {
        console.error('local storage is not working');
      }
    }
  }, [user, navigate]);


  return (
    <AuthForm
      type={ formName }
      form={ form }
      onChange={ onChange }
      onSubmit={ onSubmit }
      error={ error }
    />
  );
};

export default LoginForm;