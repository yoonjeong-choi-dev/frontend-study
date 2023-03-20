import ConfirmModal from '../../common/ConfirmModal';

const ConfirmRemoveModal = ({ visible, onConfirm, onCancel }) => {
  return (
    <ConfirmModal
      visible={ visible }
      title='Delete Post'
      description='Do you really delete this post?'
      confirmText='Delete'
      onConfirm={ onConfirm }
      onCancel={ onCancel }
    />
  );
};

export default ConfirmRemoveModal;