import router from "./router/index";
import { Switch, Route } from "react-router-dom";

const App = () => {
  const appname = `errntry`;

  return (
    <div className="w-full min-h-screen">
      <Switch>
        {router.map((route, index) => (
          <Route
            path={route.path}
            exact={route.isExact}
            render={() => <route.component appname={appname} />}
            key={index}
          />
        ))}
      </Switch>
    </div>
  );
};

export default App;
