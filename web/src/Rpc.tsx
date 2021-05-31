import { WatcherServiceClient } from "src/api/WatcherServiceClientPb";
import { SubscriptionServiceClient } from "./api/SubscriptionServiceClientPb";

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

export const subscriptionService = new SubscriptionServiceClient(
  gitpodExposedURL(),
  null,
  null
);
