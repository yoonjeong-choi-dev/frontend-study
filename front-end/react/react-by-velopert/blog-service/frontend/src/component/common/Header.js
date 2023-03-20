import styled from 'styled-components';
import { Link } from 'react-router-dom';
import Responsive from './Responsive';
import Button from './Button';

const SCWrapper = styled.div`
  position: fixed; // 상단 유지
  width: 100%;
  background: white;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.08);
`;

const SCResponsiveWrapper = styled(Responsive)`
  height: 4rem;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .logo {
    font-size: 1.125rem;
    font-weight: 800;
    letter-spacing: 2px;
  }

  .right {
    display: flex;
    align-items: center;
  }
`;

const SCUserInfo = styled.div`
  font-weight: 8000;
  margin-right: 1rem;
`;

const SCSpacer = styled.div`
  height: 4rem;
`;

const Header = ({ user, onLogout }) => {
  return (
    <>
      <SCWrapper>
        <SCResponsiveWrapper>
          <Link to='/' className='logo'>Blog Service</Link>
          { user ? (
            <div className='right'>
              <SCUserInfo>{ user.name }</SCUserInfo>
              <Button onClick={ onLogout }>Sign Out</Button>
            </div>
          ) : (
            <div className='right'>
              <Button to='/login'>Sign In</Button>
            </div>
          ) }
        </SCResponsiveWrapper>
      </SCWrapper>
      <SCSpacer />
    </>
  );
};

export default Header;