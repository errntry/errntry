import Home from "../views/Home";
import About from "../views/About";

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
  }
];

export default router;
