import { useDispatch, useSelector } from 'react-redux';
import { useCallback, useEffect } from 'react';
import { changeField, initializeField } from '../../redux/post/write';
import Editor from '../../component/post/write/Editor';

const EditorContainer = () => {
  const dispatch = useDispatch();
  const { title, body } = useSelector((state) => ({
    title: state.write.title,
    body: state.write.body,
  }));

  const onChangeField = useCallback((payload) =>
      dispatch(changeField(payload))
    , [dispatch]);

  useEffect(() => {
    // initialize the form state when unmount
    return () => {
      dispatch(initializeField());
    }
  },[]);

  return <Editor onChangeField={onChangeField} title={title} body={body}/>
};

export default EditorContainer;