import UsersContainer from '../container/UsersContainer';
import { Route, Routes } from 'react-router-dom';
import UserPage from './UserPage';

const UsersPage = () => (
  <div style={{padding: '10px 0 0 20px', display: 'flex', gap: '1rem'}}>
    <UsersContainer/>
    <Routes>
      <Route path={':id'} element={<UserPage/>}/>
    </Routes>
  </div>
);

export default UsersPage;