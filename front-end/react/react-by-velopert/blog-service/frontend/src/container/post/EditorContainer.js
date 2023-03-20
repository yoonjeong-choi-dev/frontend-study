import { useDispatch, useSelector } from 'react-redux';
import { useCallback, useEffect } from 'react';
import { changeField, initializeField } from '../../redux/post/write';
import Editor from '../../component/post/write/Editor';

const EditorContainer = ({isEdit}) => {
  const dispatch = useDispatch();
  const { title, body } = useSelector((state) => ({
    title: state.write.title,
    body: state.write.body,
  }));

  const onChangeField = useCallback((payload) =>
      dispatch(changeField(payload))
    , [dispatch]);

  // TODO
  useEffect(() => {
    if(!isEdit) {
      dispatch(initializeField());
    }
  }, [isEdit]);

  return <Editor onChangeField={ onChangeField } title={ title } body={ body } />;
};

export default EditorContainer;