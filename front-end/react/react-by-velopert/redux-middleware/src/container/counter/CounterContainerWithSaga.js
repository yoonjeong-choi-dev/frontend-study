import { connect } from 'react-redux';

import { increaseActionCreatorForSaga, decreaseActionCreatorForSaga } from '../../redux/counter';

const CounterContainerWithSaga = ({increaseAsync, decreaseAsync}) => {
  return (
    <div style={{display: 'flex', gap: '1rem', marginTop: '10px'}}>
      <span><strong>saga</strong> Dispatch with 1s async :</span>
      <button onClick={increaseAsync}>async +1</button>
      <button onClick={decreaseAsync}>async -1</button>
    </div>
  );
}

export default connect(
  null,
  {
    increaseAsync: increaseActionCreatorForSaga,
    decreaseAsync: decreaseActionCreatorForSaga}
)(CounterContainerWithSaga);