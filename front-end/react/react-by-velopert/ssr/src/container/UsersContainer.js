import { useDispatch, useSelector } from 'react-redux';
import { useEffect } from 'react';
import { getUsers } from '../redux/users';
import Users from '../component/Users';
import { Preloader, usePreloader } from '../lib/PreloadContext';

const UsersContainer = () => {
  const users = useSelector(state => state.users.users);
  const dispatch = useDispatch();

  usePreloader(() => dispatch(getUsers()));

  useEffect(() => {
    // 스토어에 데이터가 없는 경우에만 다시 호출
    // => 서버 사이드 렌더링 후에 다시 브라우저에서 호출되는 상황을 막기 위함
    if (users) return;
    dispatch(getUsers());
  }, [dispatch, users]);

  return (
    <div style={{flexGrow: '1'}}>
      <h2>Users</h2>
      <Users users={users}/>
      {/*<Preloader resolve={() => dispatch(getUsers())}/>*/}
    </div>
  );
};

export default UsersContainer;