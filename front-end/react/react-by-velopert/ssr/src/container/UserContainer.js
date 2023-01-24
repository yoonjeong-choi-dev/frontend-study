import { useDispatch, useSelector } from 'react-redux';
import { useEffect } from 'react';
import { getUser } from '../redux/users';
import { Preloader } from '../lib/PreloadContext';
import User from '../component/User';

const UserContainer = ({id}) => {
  const user = useSelector(state => state.users.user);
  const dispatch = useDispatch();

  useEffect(() => {
    // 스토어에 데이터가 없는 경우에만 다시 호출
    // => 서버 사이드 렌더링 후에 다시 브라우저에서 호출되는 상황을 막기 위함
    if (user && user.id === parseInt(id, 10)) return;
    dispatch(getUser(id));
  }, [dispatch, id, user]);

  if (!user) {
    return <Preloader resolve={() => dispatch(getUser(id))}/>;
  }
  return (
    <div style={{flexGrow: '1', border: 'solid 1px blue', paddingLeft: '10px'}}>
      <h2>User Detail - {user.name}</h2>
      <User user={user}/>
    </div>
  );
}

export default UserContainer;