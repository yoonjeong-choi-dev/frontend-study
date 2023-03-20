import styled from 'styled-components';
import Button from './Button';

const SCWrapper = styled.div`
  position: fixed;
  z-index: 30;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.25);
  display: flex;
  justify-content: center;
  align-items: center;
`;

const SCModalWrapper = styled.div`
  width: 320px;
  background: white;
  padding: 1.5rem;
  border-radius: 4px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.125);

  h2 {
    margin-top: 0;
    margin-bottom: 1rem;
  }

  p {
    margin-bottom: 3rem;
  }

  .buttons {
    display: flex;
    justify-content: flex-end;
  }
`;

const SCButton = styled(Button)`
  height: 2rem;

  & + & {
    margin-left: 0.75rem;
  }
`;

const ConfirmModal = ({
  visible,
  title,
  description,
  confirmText = 'Confirm',
  onConfirm,
  cancelText = 'Cancel',
  onCancel,
}) => {
  if(!visible) return;

  return (
    <SCWrapper>
      <SCModalWrapper>
        <h2>{title}</h2>
        <p>{description}</p>
        <div className="buttons">
          <SCButton onClick={onCancel}>{cancelText}</SCButton>
          <SCButton indigo onClick={onConfirm}>{confirmText}</SCButton>
        </div>
      </SCModalWrapper>
    </SCWrapper>
  );
};

export default ConfirmModal;