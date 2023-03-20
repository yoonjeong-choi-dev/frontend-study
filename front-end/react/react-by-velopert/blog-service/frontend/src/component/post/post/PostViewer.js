import styled from 'styled-components';

import palette from '../../../lib/styles/palette';
import Responsive from '../../common/Responsive';
import SubInfo from '../../common/SubInfo';
import Tags from '../../common/Tags';

const SCWrapper = styled(Responsive)`
  margin-top: 4rem;
`;

const SCHeader = styled.div`
  border-bottom: 1px solid ${ palette.gray[2] };
  padding-bottom: 3rem;
  margin-bottom: 3rem;

  h1 {
    font-size: 3rem;
    line-height: 1.5;
    margin: 0;
  }
`;

const SCContent = styled.div`
  font-size: 1.3125rem;
  color: ${ palette.gray[8] };;
`;

const PostViewer = ({ post, error, loading, actionButtons }) => {
  if (error) {
    if (error.response) {
      const code = error.response.status;
      if (code === 404)
        return <SCWrapper>존재하지 않는 게시글입니다</SCWrapper>;
      if (code === 400)
        return <SCWrapper>잘못된 게시글 아이디입니다</SCWrapper>;
    }
    return <SCWrapper>Error Occurred!</SCWrapper>;
  }

  if (loading || !post) {
    return null;
  }

  const { title, body, user, createdAt, tags } = post;

  return (
    <SCWrapper>
      <SCHeader>
        <h1>{ title }</h1>
        <SubInfo
          name={ user.name }
          createdAt={ createdAt }
          hasMarginTop={ true }
        />
        <Tags tags={ tags } />
      </SCHeader>
      { actionButtons }
      <SCContent dangerouslySetInnerHTML={ { __html: body } } />
    </SCWrapper>
  );
};

export default PostViewer;

