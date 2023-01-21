import { createContext, useContext } from 'react';

// 서버 사이드 렌더링 시에만 필요한 작업
// ex) 서버 사이드 렌더링 시, 비동기 작업의 결과를 리덕스 스토어에 저장해야 함
const PreloadContext = createContext(null);
export default PreloadContext;

export const Preloader = ({resolve}) => {
  const preloadContext = useContext(PreloadContext);

  // 서버 사이드 측이 아니거나, 이미 작업이 완료(서버 사이드 측에서 스토어에 저장)한 경우에는 추가적인 작업 X
  if (!preloadContext) return null;
  if (preloadContext.done) return null;

  // 서버 측에서 해당 컴포넌트가 있는 페이지를 렌더링(ssr)하는 경우에만 호출
  preloadContext.promises.push(Promise.resolve(resolve()));
  return null;
}