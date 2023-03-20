import { useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';

import { POST_ACTION_TYPE, readPost, unloadPost } from '../../redux/post/post';
import PostViewer from '../../component/post/post/PostViewer';
import ActionButtons from '../../component/post/post/ActionButtons';
import { setOriginalPost } from '../../redux/post/write';
import { remove } from '../../lib/api/post';

const PostViewerContainer = () => {
  const { postId } = useParams();
  const navigate = useNavigate();
  const dispatch = useDispatch();
  const { post, error, loading, user } = useSelector((state) => ({
    post: state.post.post,
    error: state.post.error,
    loading: state.loading[POST_ACTION_TYPE],
    user: state.user.user,
  }));

  useEffect(() => {
    dispatch(readPost(postId));
    return () => {
      dispatch(unloadPost());
    };
  }, [dispatch, postId]);

  const handleEditPost = () => {
    dispatch(setOriginalPost(post));
    navigate('/edit');
  };

  const handleRemovePost = async () => {
    try {
      await remove(postId);
      navigate('/');
    } catch (e) {
      console.error(e);
    }
  };

  const isOwner = (user && user._id) === (post && post.user._id);

  return (
    <PostViewer
      post={ post }
      loading={ loading }
      error={ error }
      actionButtons={ isOwner && (
        <ActionButtons
          onEdit={ handleEditPost }
          onRemove={ handleRemovePost }
        />
      ) }
    />
  );
};

export default PostViewerContainer;