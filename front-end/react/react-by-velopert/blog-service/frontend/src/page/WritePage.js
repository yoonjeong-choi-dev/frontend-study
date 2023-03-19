import Responsive from '../component/common/Responsive';
import EditorContainer from '../container/post/EditorContainer';
import TagBoxContainer from '../container/post/TagBoxContainer';
import ActionButtonsContainer from '../container/post/ActionButtonsContainer';


const WritePage = () => {
  return (
    <Responsive>
      <EditorContainer />
      <TagBoxContainer />
      <ActionButtonsContainer />
    </Responsive>
  );
};

export default WritePage;