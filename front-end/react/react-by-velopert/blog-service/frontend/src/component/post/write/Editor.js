import styled from 'styled-components';
import Quill from 'quill';
import 'quill/dist/quill.bubble.css';

import { useEffect, useRef } from 'react';

import Responsive from '../../common/Responsive';
import palette from '../../../lib/styles/palette';

const SCWrapper = styled(Responsive)`
  padding: 5rem 0;
`;

const SCTitleInput = styled.input`
  font-size: 3rem;
  outline: none;
  width: 100%;
  margin-bottom: 2rem;
  padding-bottom: 0.5rem;
  border: none;
  border-bottom: 1px solid ${ palette.gray[4] };
`;

const SCQuillWrapper = styled.div`
  .ql-editor {
    padding: 0;
    min-height: 320px;
    font-size: 1.125rem;
    line-height: 1.5;
  }

  .ql-editor.ql-blank:before {
    left: 0;
  }
`;

const Editor = ({ onChangeField, body, title }) => {
  const quillElem = useRef(null);
  const quillInst = useRef(null);

  useEffect(() => {
    quillInst.current = new Quill(quillElem.current, {
      theme: 'bubble',
      placeholder: 'Enter contents...',
      modules: {
        toolbar: [
          [{ header: '1' }, { header: '2' }],
          ['bold', 'italic', 'underline', 'strike'],
          [{ list: 'ordered' }, { list: 'bullet' }],
          ['blockquote', 'code-block', 'link', 'image'],
        ],
      },
    });

    const quill = quillInst.current;
    quill.on('text-change', (delta, oldDelta, source) => {
      if (source === 'user') {
        onChangeField({ key: 'body', value: quill.root.innerHTML });
      }
    });
  }, [onChangeField]);

  const isMounted = useRef(false);
  useEffect(() => {
    if (isMounted.current) return;

    isMounted.current = true;
    quillInst.current.root.innerHTML = body;
  }, [body]);

  const handleChangeTitle = (e) => {
    onChangeField({ key: 'title', value: e.target.value });
  };

  return (
    <SCWrapper>
      <SCTitleInput onChange={ handleChangeTitle } value={ title } placeholder='Enter title...' />
      <SCQuillWrapper>
        <div ref={ quillElem } />
      </SCQuillWrapper>
    </SCWrapper>
  );
};

export default Editor;