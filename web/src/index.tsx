import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import Paperbase from "./Paperbase";
import reportWebVitals from "./reportWebVitals";

ReactDOM.render(
  <React.StrictMode>
    <App />
    <Paperbase />
  </React.StrictMode>,
  document.getElementById("root")
);

reportWebVitals();
