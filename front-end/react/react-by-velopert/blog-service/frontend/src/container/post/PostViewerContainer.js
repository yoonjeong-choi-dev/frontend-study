import { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';

import { POST_ACTION_TYPE, readPost, unloadPost } from '../../redux/post/post';
import PostViewer from '../../component/post/post/PostViewer';

const PostViewerContainer = () => {
  const { postId } = useParams();
  const dispatch = useDispatch();
  const { post, error, loading } = useSelector((state) => ({
    post: state.post.post,
    error: state.post.error,
    loading: state.loading[POST_ACTION_TYPE],
  }));

  useEffect(() => {
    dispatch(readPost(postId));
    return () => {
      dispatch(unloadPost());
    };
  }, [dispatch, postId]);

  return (
    <PostViewer
      post={ post }
      loading={ loading }
      error={ error }
    />
  );
};

export default PostViewerContainer;