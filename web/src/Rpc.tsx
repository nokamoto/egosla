import { WatcherServiceClient } from "./api/WatcherServiceClientPb";

function gitpodExposedURL(): string {
  return (
    window.location.protocol +
    "//" +
    window.location.hostname.replace("3000", "8080")
  );
}

export const watcherService = new WatcherServiceClient(
  gitpodExposedURL(),
  null,
  null
);
