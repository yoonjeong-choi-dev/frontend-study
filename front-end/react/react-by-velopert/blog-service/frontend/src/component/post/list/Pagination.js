import styled from 'styled-components';
import qs from 'qs';
import Button from '../../common/Button';

const SCWrapper = styled.div`
  width: 320px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  margin-bottom: 3rem;
`;

const SCPageIndexWrapper = styled.div``;

const buildLink = ({ username, tag, page }) => {
  const query = qs.stringify({ tag, page });
  return username ? `/${ username }?${ query }` : `/?${ query }`;
};

const Pagination = ({ page, lastPage, username, tag }) => {
  return (
    <SCWrapper>
      <Button
        disabled={ page === 1 }
        to={ page === 1 ? undefined : buildLink({ username, tag, page: page - 1 }) }
      >
        Prev
      </Button>
      <SCPageIndexWrapper>{ page }</SCPageIndexWrapper>
      <Button
        disabled={ page === lastPage }
        to={ page === lastPage ? undefined : buildLink({ username, tag, page: page + 1 }) }
      >
        Next
      </Button>
    </SCWrapper>
  );
};

export default Pagination;