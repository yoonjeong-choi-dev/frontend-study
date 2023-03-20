import { Helmet } from 'react-helmet-async';

import Responsive from '../component/common/Responsive';
import EditorContainer from '../container/post/EditorContainer';
import TagBoxContainer from '../container/post/TagBoxContainer';
import ActionButtonsContainer from '../container/post/ActionButtonsContainer';
import React from 'react';


const WritePage = () => {
  return (
    <Responsive>
      <Helmet>
        <title>Blog Service - Write Post</title>
      </Helmet>
      <EditorContainer isEdit={false} />
      <TagBoxContainer />
      <ActionButtonsContainer />
    </Responsive>
  );
};

export default WritePage;