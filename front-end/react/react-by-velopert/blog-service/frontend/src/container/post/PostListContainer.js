import { useParams, useSearchParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { listPosts, POST_LIST_ACTION_TYPE } from '../../redux/post/posts';
import { useEffect } from 'react';
import PostList from '../../component/post/list/PostList';

const PostListContainer = () => {
  const { username } = useParams();
  const [searchParams] = useSearchParams();

  const dispatch = useDispatch();
  const { posts, error, loading, user } = useSelector((state) => ({
    posts: state.posts.posts,
    error: state.posts.error,
    loading: state.loading[POST_LIST_ACTION_TYPE],
    user: state.user.user,
  }));

  useEffect(() => {
    const tag = searchParams.get('tag');
    const page = parseInt(searchParams.get('page'), 10) || 1;
    dispatch(listPosts({ username, tag, page }));
  }, [dispatch, searchParams, username]);

  return (
    <PostList
      posts={ posts }
      loading={ loading }
      error={ error }
      user={ user }
    />
  );
};

export default PostListContainer;