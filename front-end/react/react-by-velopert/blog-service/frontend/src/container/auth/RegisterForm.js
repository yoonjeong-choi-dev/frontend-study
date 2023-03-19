import { useDispatch, useSelector } from 'react-redux';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { changeField, initializeForm, register } from '../../redux/auth/auth';
import AuthForm from '../../component/auth/AuthForm';
import { check } from '../../redux/auth/user';


const formName = 'register';

const RegisterForm = () => {
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

    const { name, password, passwordConfirm } = form;

    if ([name, password, passwordConfirm].includes('')) {
      setError('모든 정보를 입력해주세요');
      return;
    }

    if (password !== passwordConfirm) {
      setError('비밀번호가 일치하지 않습니다');
      return;
    }

    dispatch(register({ name, password }));
  };

  useEffect(() => {
    dispatch(initializeForm(formName));
  }, [dispatch]);


  useEffect(() => {
    if (authError) {
      console.log('Error for register: ', authError);
      const code = authError.response.status;
      if (code === 409) {
        setError('이미 존재하는 회원입니다');
      } else {
        setError('회원가입 실패');
        return;
      }
    }
    if (auth) {
      console.log('Success to register: ', auth);
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

export default RegisterForm;