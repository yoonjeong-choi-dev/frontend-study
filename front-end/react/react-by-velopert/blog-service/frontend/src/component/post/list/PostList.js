import styled from 'styled-components';

import palette from '../../../lib/styles/palette';
import Responsive from '../../common/Responsive';
import Button from '../../common/Button';
import SubInfo from '../../common/SubInfo';
import Tags from '../../common/Tags';
import { Link } from 'react-router-dom';

const SCWrapper = styled(Responsive)`
  margin-top: 3rem;
`;

const SCWriteButtonWrapper = styled.div`
  display: flex;
  justify-content: flex-end;
  margin-bottom: 3rem;
`;

const SCItemWrapper = styled.div`
  padding: 3rem 0;

  &:first-child {
    padding-top: 0;
  }

  & + & {
    border-top: 1px solid ${ palette.gray[2] };
  }

  h2 {
    font-size: 2rem;
    margin-top: 0;
    margin-bottom: 0;

    &:hover {
      color: ${ palette.gray[6] };
    }
  }

  p {
    margin-top: 2rem;
  }
`;

const PostItem = ({ post }) => {
  const { user, tags, title, body, _id, createdAt } = post;
  return (
    <SCItemWrapper>
      <h2>
        <Link to={ `/${ user.name }/${ _id }` }>
          { title }
        </Link>
      </h2>
      <SubInfo name={ user.name } createdAt={ createdAt } />
      <Tags tags={ tags } />
      <p>{ body }</p>
    </SCItemWrapper>
  );
};

const PostList = ({ posts, loading, error, user }) => {
  if (error) {
    if (error.response) {
      const code = error.response.status;
      if (code === 400)
        return <SCWrapper>잘못된 요청입니다</SCWrapper>;
    }
    return <SCWrapper>Error Occurred!</SCWrapper>;
  }

  return (
    <SCWrapper>
      <SCWriteButtonWrapper>
        { user && (
          <Button indigo to='/write'>
            Write a Post
          </Button>
        ) }
      </SCWriteButtonWrapper>
      { !loading && posts && (
        <div>
          { posts.map((post) => (
            <PostItem post={ post } key={ post._id } />
          )) }
        </div>
      ) }
    </SCWrapper>
  );
};

export default PostList;