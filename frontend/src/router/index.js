import Home from "../views/Home";
import About from "../views/About";
import Signup from "../views/Signup";
import Login from "../views/Login";

const router = [
  {
    path: "/",
    name: "Home",
    component: Home,
    isExact: true
  },
  {
    path: "/about",
    name: "About",
    component: About,
    isExact: false
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
    isExact: false
  },
  {
    path: "/signup",
    name: "Signup",
    component: Signup,
    isExact: false
  }
];

export default router;
