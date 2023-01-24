import { connect } from 'react-redux';

import { decrease, increase } from '../../redux/counter';
import Counter from '../../component/Counter';

const CounterContainer = ({number, increase, decrease}) => {
  return (
    <Counter number={number} onIncrease={increase} onDecrease={decrease}/>
  );
}

export default connect(
  state => ({
    number: state.counter,
  }),
  {increase, decrease}
)(CounterContainer);