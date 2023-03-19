import { useDispatch, useSelector } from 'react-redux';
import { changeField } from '../../redux/post/write';
import TagBox from '../../component/post/write/TagBox';

const TagBoxContainer = () => {
  const dispatch = useDispatch();
  const tags = useSelector((state) => state.write.tags);

  const handleChangeTags = (newTags) => {
    dispatch(changeField({
      key: 'tags',
      value: newTags,
    }));
  };

  return <TagBox onChangeTags={ handleChangeTags } tags={ tags } />;
};

export default TagBoxContainer;