import { useState } from "react";
import PropTypes from "prop-types";
import image from "../assets/bug.png";
import { Link } from "react-router-dom";

const Home = ({ appname }) => {
  const [trial, setTrial] = useState(true);

  return (
    <div className="">
      <div
        className="flex flex-col items-center justify-center w-full h-full px-5 py-10 text-white md:justify-between md:flex-row bg-gradient-to-tr from-gray-900 to-cyan-500 animate-moveBackground xs:px-7 sm:px-12 md:px-24 lg:px-36"
        style={{
          backgroundSize: "400% 400%",
          clipPath: "polygon(0 0, 100% 0%, 100% 95%, 0% 100%)"
        }}
      >
        <div className="w-full mx-auto md:w-1/2 md:block">
          <img src={image} alt="404" className="w-full h-full" />
        </div>
        <div className="w-auto h-full px-4 space-y-10 text-center md:text-justify">
          <div className="text-3xl font-bold uppercase lg:text-4xl">
            {appname}
          </div>
          <div className="">
            Errntry's application monitoring platform helps every developer
            diagnose, fix, and optimize the performance of their code.
          </div>

          {trial ? (
            <button
              type="button"
              onClick={() => setTrial(!trial)}
              className={`${
                trial ? `block` : `hidden`
              } w-full py-2 mx-auto uppercase border md:w-1/2 lg:w-1/3 shadow-sm transform hover:-translate-y-0.5 transition-all duration-200 font-semibold rounded-xl focus:outline-none  hover:bg-white hover:text-cyan-800`}
            >
              Try {appname}
            </button>
          ) : (
            <div className={`flex justify-around`}>
              <Link to="/login">
                <button
                  type="button"
                  className={`px-5 border border-white py-1.5 uppercase font-medium rounded-sm hover:bg-white hover:text-cyan-800 shadow-sm transform hover:-translate-y-0.5 transition-all duration-200`}
                >
                  Login
                </button>
              </Link>
              <Link to="/signup">
                <button
                  type="button"
                  className={`px-5 border border-white py-1.5 uppercase font-medium rounded-sm hover:bg-white hover:text-cyan-800 shadow-sm transform hover:-translate-y-0.5 transition-all duration-200`}
                >
                  Signup
                </button>
              </Link>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

Home.propTypes = {
  appname: PropTypes.string.isRequired
};

export default Home;
