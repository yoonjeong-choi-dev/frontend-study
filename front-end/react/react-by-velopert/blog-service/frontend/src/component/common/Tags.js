import styled from 'styled-components';
import palette from '../../lib/styles/palette';
import { Link } from 'react-router-dom';

const SCWrapper = styled.div`
  margin-top: 0.5rem;

  .tag {
    display: inline-block;
    color: ${ palette.indigo[7] };
    text-decoration: none;
    margin-right: 0.5rem;

    &:hover {
      color: ${ palette.indigo[6] };
    }
  }
`;

const Tags = ({ tags }) => {
  return (
    <SCWrapper>
      { tags.map((tag) => (
        <div className='tag' key={tag}>
          <Link to={ `/?tag=${ tag }` }>
            #{ tag }
          </Link>
        </div>
      )) }
    </SCWrapper>
  );
};

export default Tags;