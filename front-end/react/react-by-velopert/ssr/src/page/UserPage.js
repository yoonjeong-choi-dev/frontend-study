import UserContainer from '../container/UserContainer';
import { useParams } from 'react-router-dom';

const UserPage = () => {
  const {id} = useParams();
  return (
    <div style={{flexGrow: '2', margin: '10px'}}>
      <UserContainer id={id}/>
    </div>
  );
};

export default UserPage;