import HeaderContainer from '../container/common/HeaderContainer';
import PostListContainer from '../container/post/PostListContainer';
import PaginationContainer from '../container/post/PaginationContainer';

const PostListPage = () => {
  return (
    <>
      <HeaderContainer />
      <PostListContainer />
      <PaginationContainer />
    </>
  );
};

export default PostListPage;