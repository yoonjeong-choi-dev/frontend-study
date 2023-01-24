import SimpleInfo from './SimpleInfo';

const User = ({user}) => {
  const {name, username, email, phone, website} = user;

  return (
    <div>
      <h3>User : {username}({name})</h3>
      <SimpleInfo title="email" content={email}/>
      <SimpleInfo title="phone" content={phone}/>
      <SimpleInfo title="website" content={website}/>
    </div>
  );
}

export default User;