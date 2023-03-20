import { finishLoading, startLoading } from '../redux/loading';

export const createRequestActionTypes = (type) => {
  const SUCCESS = `${ type }_SUCCESS`;
  const FAILURE = `${ type }_FAILURE`;
  return [type, SUCCESS, FAILURE];
};

export default function createRequestThunk(type, request) {
  const SUCCESS = `${ type }_SUCCESS`;
  const FAILURE = `${ type }_FAILURE`;

  // redux-thunk 미들웨어가 처리할 thunk 함수 반환
  return (params) => async (dispatch) => {
    // 요청 시작 -> 로딩
    dispatch({ type });
    dispatch(startLoading(type));

    // API call
    try {
      const response = await request(params);
      dispatch({
        type: SUCCESS,
        payload: response.data,
        meta: response,
      });
    } catch (e) {
      dispatch({
        type: FAILURE,
        payload: e,
        error: true,
      });
    } finally {
      dispatch(finishLoading(type));
    }
  };
}