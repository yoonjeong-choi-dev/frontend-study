import { combineReducers } from 'redux';

import counter from './counter';
import todos from './todos';

// 스토어 생성 시 호출하는 createStore 함수는 하나의 리듀서만 사용 가능
// => combineReducers 로 여러 리듀서를 하나로 묶어준다
const rootReducer = combineReducers({
  counter,
  todos
});

export default rootReducer;