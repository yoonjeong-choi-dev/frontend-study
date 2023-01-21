import { Link } from 'react-router-dom';

const Menu = () => (
  <ul>
    <li>
      <Link to={'/'}>Home</Link>
    </li>
    <li>
      <Link to={'/red'}>Red</Link>
    </li>
    <li>
      <Link to={'/blue'}>Blue</Link>
    </li>
  </ul>
);

export default Menu;