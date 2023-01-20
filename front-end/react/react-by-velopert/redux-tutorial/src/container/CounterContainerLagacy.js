import { connect } from 'react-redux';

import Counter from '../component/Counter';
import { increase, decrease } from '../redux/counter';

const CounterContainer = ({number, increase, decrease}) => {
  return <Counter number={number} onIncrease={increase} onDecrease={decrease}/>;
}

// mapStateToProps : 리덕스 스토어에 있는 상태를 컴포넌트의 props로 넘겨주기 위한 함수
const mapStateToProps = state => ({
  number: state.counter.number
});

// mapDispatchToProps : 액션 생성 함수가 생성하는 액션을 컴포넌트의 props로 넘겨주기 위한 함수
const mapDispatchToProps = dispatch => ({
  increase: () => {
    dispatch(increase());
  },
  decrease: () => {
    dispatch(decrease());
  },
});

export default connect(mapStateToProps, mapDispatchToProps)(CounterContainer);