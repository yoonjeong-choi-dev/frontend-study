import { useEffect } from 'react';
import { connect } from 'react-redux';

import JsonHolder from '../component/JsonHolder';
import { GET_POST, GET_USERS, getPost, getUsers } from '../redux/jsonHolderForSaga';

const JsonHolderContainerForSaga = (
  {
    getPost,
    getUsers,
    post,
    users,
    loadingPost,
    loadingUsers
  }) => {
  useEffect(() => {
    getPost(1);
    getUsers();
  }, [getPost, getUsers]);

  return (
    <JsonHolder
      title={"Saga Example"}
      post={post}
      users={users}
      loadingPost={loadingPost}
      loadingUsers={loadingUsers}
    />
  );
}

export default connect(
  ({jsonHolder, loading}) => ({
    loadingPost: loading[GET_POST],
    loadingUsers: loading[GET_USERS],
    post: jsonHolder.post,
    users: jsonHolder.users,
  }),
  {getPost, getUsers}
)(JsonHolderContainerForSaga);