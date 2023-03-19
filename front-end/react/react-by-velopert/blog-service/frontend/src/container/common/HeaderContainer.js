import { useDispatch, useSelector } from 'react-redux';
import Header from '../../component/common/Header';
import { logout } from '../../redux/user';

const HeaderContainer = () => {
  const { user } = useSelector((state) => ({ user: state.user.user }));
  const dispatch = useDispatch();
  const handleLogout = () => {
    dispatch(logout());
  };
  return <Header user={ user } onLogout={ handleLogout } />;
};

export default HeaderContainer;