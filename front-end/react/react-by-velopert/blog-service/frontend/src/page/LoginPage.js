import AuthTemplate from '../component/auth/AuthTemplate';
import LoginForm from '../container/auth/LoginForm';

const LoginPage = () => {
  return (
    <AuthTemplate>
      <LoginForm />
    </AuthTemplate>
  );
};

export default LoginPage;