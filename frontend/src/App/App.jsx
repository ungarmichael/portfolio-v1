import React from "react";
import { Route, Switch } from "react-router-dom";

import About from "./pages/About/about.js";
import Projects from "./pages/Projects/projects.js";
import Content from "./pages/Content/content.js";
import Contact from "./pages/Contact/contact.js";
import Error from "./pages/Error/error.js";
import Foyer from "./pages/Foyer/foyer.js";
import UgNavbar from './shared/navbar/navbar';

function App() {
  return(
    <React.Fragment>
      <UgNavbar/>
      <Switch>
      <Route exact path="/" component={Foyer}></Route>
      <Route path="/Projects" component={Projects}></Route>
      <Route path="/Content" component={Content}></Route>
      <Route path="/Contact" component={Contact}></Route>
      <Route path="/About" component={About}></Route>
      <Route component={Error}></Route>
      </Switch>
    </React.Fragment>
  );
}

export default App;
