import styled from 'styled-components';

import palette from '../../../lib/styles/palette';
import { useState } from 'react';
import ConfirmRemoveModal from './ConfirmRemoveModal';

const SCWrapper = styled.div`
  display: flex;
  justify-content: flex-end;
  margin: -1.5rem 0 2rem;
`;

const SCActionButton = styled.button`
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
  color: ${ palette.gray[6] };
  font-weight: bold;
  border: none;
  outline: none;
  font-size: 0.875rem;
  cursor: pointer;

  &:hover {
    background: ${ palette.gray[1] };
    color: ${ palette.indigo[7] };
  }

  & + & {
    margin-left: 0.25rem;
  }
`;

const ActionButtons = ({ onEdit, onRemove }) => {
  const [modal, setModal] = useState(false);
  const handleRemoveButtonClick = () => {
    setModal(true);
  };
  const handleCancelModal = () => {
    setModal(false);
  };

  const handleConfirmModal = () => {
    setModal(false);
    onRemove();
  };

  return (
    <>
      <SCWrapper>
        <SCActionButton onClick={ onEdit }>Edit</SCActionButton>
        <SCActionButton onClick={ handleRemoveButtonClick }>Delete</SCActionButton>
      </SCWrapper>
      <ConfirmRemoveModal
        visible={ modal }
        onConfirm={ handleConfirmModal }
        onCancel={ handleCancelModal }
      />
    </>
  );
};

export default ActionButtons;