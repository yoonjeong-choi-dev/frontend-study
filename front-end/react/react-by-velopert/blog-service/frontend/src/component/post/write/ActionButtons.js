import styled from 'styled-components';
import Button from '../../common/Button';

const SCWrapper = styled.div`
  margin: 1rem 0 3rem;

  button + button {
    margin-left: 0.5rem;
  }
`;

const SCButton = styled(Button)`
  height: 2.125rem;

  & + & {
    margin-left: 0.5rem;
  }
`;

const SCErrorMessage = styled.div`
  color: red;
  font-size: 0.875rem;
  margin-top: 1rem;
`;

const ActionButtons = ({ onCancel, onSubmit, error, isEdit }) => {
  return (
    <SCWrapper>
      <SCButton indigo onClick={ onSubmit }>
        { isEdit ? 'Edit' : 'Create' }
      </SCButton>
      <SCButton onClick={ onCancel }>
        Cancel
      </SCButton>
      { error && (
        <SCErrorMessage>
          { error }
        </SCErrorMessage>
      ) }
    </SCWrapper>
  );
};

export default ActionButtons;