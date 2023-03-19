import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { writePost } from '../../redux/post/write';
import { useEffect, useState } from 'react';
import ActionButtons from '../../component/post/write/ActionButtons';


const ActionButtonsContainer = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const { title, body, tags, post, postError } = useSelector((state) => ({
    title: state.write.title,
    body: state.write.body,
    tags: state.write.tags,
    post: state.write.post,
    postError: state.write.postError,
  }));

  const [error, setError] = useState(null);

  const handleCreatePost = () => {
    if ([title, body, tags].includes('')) {
      setError('모든 정보를 입력해주세요');
      return;
    }

    dispatch(writePost({
      title, body, tags,
    }));
  };

  const handleCancel = () => {
    navigate(-1);
  };

  useEffect(() => {
    if (post) {
      const { _id, user } = post;
      navigate(`/${ user.username }/${ _id }`);
    }
    if (postError) {
      const code = postError.response.status;
      if (code === 401) {
        alert('로그인 필요');
        navigate('/login');
      } else {
        console.log(postError);
        setError('게시물 등록 실패');
      }
    }
  }, [navigate, post, postError]);

  return <ActionButtons
    onSubmit={ handleCreatePost }
    onCancel={ handleCancel }
    error={ error }
  />;
};

export default ActionButtonsContainer;