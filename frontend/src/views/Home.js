import React from "react";
import PropTypes from "prop-types";

const Home = ({ appname }) => {
  return (
    <div className="w-full h-full">
      <div className="px-4 sm:px-6 lg:px-16">{appname}</div>
    </div>
  );
};

Home.propTypes = {};

export default Home;
