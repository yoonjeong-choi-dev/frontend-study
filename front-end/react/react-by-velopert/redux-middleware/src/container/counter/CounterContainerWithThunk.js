import { connect } from 'react-redux';

import { increaseThunk, decreaseThunk } from '../../redux/counter';

const CounterContainerWithThunk = ({increaseAsync, decreaseAsync}) => {
  return (
    <div style={{display: 'flex', gap: '1rem', marginTop: '10px'}}>
      <span><strong>thunk</strong> Dispatch with 1s async :</span>
      <button onClick={increaseAsync}>async +1</button>
      <button onClick={decreaseAsync}>async -1</button>
    </div>
  );
}

export default connect(
  null,
  {increaseAsync: increaseThunk, decreaseAsync: decreaseThunk}
)(CounterContainerWithThunk);