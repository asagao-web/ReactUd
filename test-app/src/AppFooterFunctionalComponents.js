import React, {Fragment} from 'react';




export default function AppFooterFunctionalComponent(props) {
  const currentYear = new Date().getFullYear();
  return (
    <Fragment>
      <hr />
      <p className="footer">Copyright &copy; {currentYear} Maya Ltd.</p>
      <p className="footer">{props.myProperty}</p>
    </Fragment>
  );
}
