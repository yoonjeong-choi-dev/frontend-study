import React, { useCallback, useEffect, useState } from 'react';
import styled from 'styled-components';
import palette from '../../../lib/styles/palette';

const SCWrapper = styled.div`
  width: 100%;
  border-top: 1px solid ${ palette.gray[2] };
  padding-top: 2rem;

  h4 {
    color: ${ palette.gray[8] };
    margin: 0 0 0.5rem;
  }
`;

const SCTagForm = styled.form`
  border-radius: 4px;
  overflow: hidden;
  display: flex;
  width: 256px;
  border: 1px solid ${ palette.gray[9] };

  input, button {
    outline: none;
    border: none;
    font-size: 1rem;
  }

  button {
    cursor: pointer;
    padding: 0 1rem;
    border: none;
    background: ${ palette.gray[8] };
    color: white;
    font-weight: bold;

    &:hover {
      background: ${ palette.gray[6] };
    }
  }

  input {
    padding: 0.5rem;
    flex: 1;
    min-width: 0;
  }
`;

const SCTag = styled.div`
  margin-right: 0.5rem;
  color: ${ palette.gray[6] };

  &:hover {
    opacity: 0.5;
  }
`;

const SCTagListWrapper = styled.div`
  display: flex;
  margin-top: 0.5rem;
`;

const TagItem = React.memo(({ tag, onRemove }) => (
  <SCTag onClick={ () => onRemove(tag) }>#{ tag }</SCTag>
));

const TagList = React.memo(({ tags, onRemove }) => (
  <SCTagListWrapper>
    { tags.map((tag) => (
      <TagItem key={ tag } tag={ tag } onRemove={ onRemove } />
    )) }
  </SCTagListWrapper>
));

const TagBox = ({ tags, onChangeTags }) => {
  const [input, setInput] = useState('');
  const [localTags, setLocalTags] = useState([]);

  useEffect(() => {
    setLocalTags(tags);
  }, [tags]);

  const onChangeInput = useCallback((e) => {
    setInput(e.target.value);
  }, []);

  const addTag = useCallback((tag) => {
    const value = tag.trim();
    if (!value) return;
    if (localTags.includes(value)) return;

    const newTags = [...localTags, value];

    setLocalTags(newTags);
    onChangeTags(newTags);
  }, [localTags, onChangeTags]);

  const removeTag = useCallback((tag) => {
    const newTags = localTags.filter(t => t !== tag);

    setLocalTags(newTags);
    onChangeTags(newTags);
  }, [localTags, onChangeTags]);

  const handleAddTag = useCallback((e) => {
    e.preventDefault();
    addTag(input);
    setInput('');
  }, [input, addTag]);

  return (
    <SCWrapper>
      <h4>Tag</h4>
      <SCTagForm onSubmit={ handleAddTag }>
        <input onChange={ onChangeInput } value={ input } placeholder='Enter a tag...' />
        <button type='submit'>Add</button>
      </SCTagForm>
      <TagList tags={ localTags } onRemove={ removeTag } />
    </SCWrapper>
  );
};

export default TagBox;