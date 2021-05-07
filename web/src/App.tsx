import React, { useEffect } from "react";
import { WatcherServiceClient } from "./api/ServiceServiceClientPb";
import { CreateWatcherRequest, Watcher } from "./api/service_pb";

function gitpodExposedURL(): string {
  return (
    window.location.protocol +
    "//" +
    window.location.hostname.replace("3000", "8080")
  );
}

function App() {
  useEffect(() => {
    const req = new CreateWatcherRequest();
    req.setWatcher(new Watcher().setKeywordsList(["foo", "bar", "baz"]));

    const service = new WatcherServiceClient(gitpodExposedURL(), null, null);
    service.createWatcher(req, {}, (err, res) => {
      console.log("err", err);
      console.log("res", res);
    });
  });
  return <div></div>;
}

export default App;
