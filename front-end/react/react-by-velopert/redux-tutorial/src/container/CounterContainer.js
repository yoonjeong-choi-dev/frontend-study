import { useDispatch, useSelector } from 'react-redux';

import Counter from '../component/Counter';
import { increase, decrease } from '../redux/counter';
import { useCallback } from 'react';

const CounterContainer = () => {
  // 기존에 정의한 mapStateToProps
  const number = useSelector(state => state.counter.number);

  // 리덕스 스토어에 dispatch 발생시키는 함수 => 인자로 액션 생성 함수를 넣어준다
  // mapDispatchToProps 에서 정의한 방식
  const dispatch = useDispatch();
  const onIncrease = useCallback(() => dispatch(increase()), []);
  const onDecrease = useCallback(() => dispatch(decrease()), []);

  return <Counter
    number={number}
    onIncrease={onIncrease}
    onDecrease={onDecrease}
  />;
}

// mapStateToProps : 리덕스 스토어에 있는 상태를 컴포넌트의 props로 넘겨주기 위한 함수
// const mapStateToProps = state => ({
//   number: state.counter.number
// });

// mapDispatchToProps : 액션 생성 함수가 생성하는 액션을 컴포넌트의 props로 넘겨주기 위한 함수
// const mapDispatchToProps = dispatch => ({
//   increase: () => {
//     dispatch(increase());
//   },
//   decrease: () => {
//     dispatch(decrease());
//   },
// });

export default CounterContainer;