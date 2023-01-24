import { finishLoading, startLoading } from '../loading';

export default function createRequestThunk(type, request) {
  const SUCCESS = `${type}_SUCCESS`;
  const FAILURE = `${type}_FAILURE`;

  // redux-thunk 미들웨어가 처리할 thunk 함수 반환
  return params => async dispatch => {
    // 요청을 시작함을 알림
    dispatch({type});
    dispatch(startLoading(type));

    // API Call
    try {
      const response = await request(params);
      dispatch({
        type: SUCCESS,
        payload: response.data,
      });
    } catch (e) {
      dispatch({
        type: FAILURE,
        payload: e,
        error: true,
      });

      throw e;
    } finally {
      dispatch(finishLoading(type));
    }
  }
}