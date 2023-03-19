import styled from 'styled-components';
import { Link } from 'react-router-dom';

import palette from '../../lib/styles/palette';
import Button from '../common/Button';

const SCFormWrapper = styled.div`
  h3 {
    margin: 0 0 1rem;
    color: ${ palette.gray[8] };
  }
`;

const SCInput = styled.input`
  font-size: 1rem;
  border: none;
  border-bottom: 1px solid ${ palette.gray[5] };
  padding-bottom: 0.5rem;
  outline: none;
  width: 100%;

  &:focus {
    color: $oc-teal-7;
    border-bottom: 1px solid ${ palette.gray[7] };
  }

  & + & {
    margin-top: 1rem;
  }
`;

const SCButtonWrapper = styled.div`
  margin-top: 1rem;
`;

const SCFooter = styled.div`
  margin-top: 2rem;
  text-align: right;

  a {
    color: ${ palette.gray[6] };
    text-decoration: underline;

    &:hover {
      color: ${ palette.gray[9] };
    }
  }
`;

const SCErrorMessage = styled.div`
  color: red;
  text-align: center;
  font-size: 0.875rem;
  margin-top: 1rem;
`;

const labelMap = {
  login: 'Sign In',
  register: 'Sign Up',
};

const AuthForm = ({ type, form, onChange, onSubmit, error }) => {
  const label = labelMap[type];

  return (
    <SCFormWrapper>
      <h3>{ label }</h3>
      <form onSubmit={ onSubmit }>
        <SCInput
          autoComplete='username'
          name='name'
          placeholder='Enter your ID'
          onChange={ onChange }
          value={ form.name }
        />
        <SCInput
          autoComplete='new-password'
          name='password'
          placeholder='Enter your password'
          type='password'
          onChange={ onChange }
          value={ form.password }
        />
        { type === 'register' && (
          <SCInput
            autoComplete='new-password'
            name='passwordConfirm'
            placeholder='Confirm your password'
            type='password'
            onChange={ onChange }
            value={ form.passwordConfirm }
          />
        ) }

        { error && (
          <SCErrorMessage>
            { error }
          </SCErrorMessage>
        ) }
        <SCButtonWrapper>
          <Button fullWidth indigo>{ label }</Button>
        </SCButtonWrapper>
      </form>
      <SCFooter>
        { type === 'register' ? (
          <Link to='/login'>{ labelMap.login }</Link>) : (
          <Link to='/register'>{ labelMap.register }</Link>) }
      </SCFooter>
    </SCFormWrapper>
  );
};

export default AuthForm;