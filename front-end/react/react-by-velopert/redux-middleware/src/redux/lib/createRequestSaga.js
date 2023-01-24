import { call, put } from 'redux-saga/effects';
import { finishLoading, startLoading } from '../loading';

export default function createRequestSaga(type, request) {
  const SUCCESS = `${type}_SUCCESS`;
  const FAILURE = `${type}_FAILURE`;

  // redux-saga 미들웨어에서 사용 하는 사가(특정 액션을 처리하는 작업)을 반환
  return function* (action) {
    // 요청 시작을 알림
    yield put(startLoading(type));

    // API Call
    try {
      const response = yield call(request, action.payload);
      yield put({
        type: SUCCESS,
        payload: response.data,
      });
    } catch (e) {
      yield put({
        type: FAILURE,
        payload: e,
        error: true,
      });

      throw e;
    } finally {
      yield put(finishLoading(type));
    }
  }
}