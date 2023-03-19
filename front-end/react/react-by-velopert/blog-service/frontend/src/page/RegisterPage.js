import AuthTemplate from '../component/auth/AuthTemplate';
import RegisterForm from '../container/auth/RegisterForm';

const RegisterPage = () => {
  return (
    <AuthTemplate>
      <RegisterForm />
    </AuthTemplate>
  );
};

export default RegisterPage;