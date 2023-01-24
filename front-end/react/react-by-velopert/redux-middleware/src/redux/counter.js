import { createAction, handleActions } from 'redux-actions';
import { delay, put, takeEvery, takeLatest } from 'redux-saga/effects';

// Action Type
const INCREASE = 'counter/INCREASE';
const DECREASE = 'counter/DECREASE';

// Action Creators
export const increase = createAction(INCREASE);
export const decrease = createAction(DECREASE);

// Redux Thunk 에서 사용할 액션 생성 함수
export const increaseThunk = () => dispatch => {
  setTimeout(() => {
    dispatch(increase());
  }, 1000);
};

export const decreaseThunk = () => dispatch => {
  setTimeout(() => {
    dispatch(decrease());
  }, 1000);
}

// Reducer
const initialState = 0;
const counter = handleActions({
  [INCREASE]: state => state + 1,
  [DECREASE]: state => state - 1,
}, initialState);

export default counter;


// Action Type for Saga
const INCREASE_ASYNC = 'counter/INCREASE_ASYNC';
const DECREASE_ASYNC = 'counter/DECREASE_ASYNC';

// Redux Sage 에서 사용하기 위한 액션 생성 함수
export const increaseActionCreatorForSaga = createAction(INCREASE_ASYNC, () => undefined);
export const decreaseActionCreatorForSaga = createAction(DECREASE_ASYNC, () => undefined);

// Redux Saga 에서 사용하는 제너레이터 함수(사가)
// => 특정 액션에서 처리할 작업
function* increaseSaga() {
  yield delay(1000);
  yield put(increase());
}

function* decreaseSaga() {
  yield delay(1000);
  yield put(decrease());
}

export function* counterSaga() {
  // 해당 액셥 타입에 대한 액션에 대해서 작업 처리
  yield takeEvery(INCREASE_ASYNC, increaseSaga);

  // 해당 액션 타입에 대해서 기존에 진행 중이던 작업이 있으면 취소 처리 후 가장 마지막 작업만 수행
  yield takeLatest(DECREASE_ASYNC, decreaseSaga);
}

