import { FieldMask } from "google-protobuf/google/protobuf/field_mask_pb";
import { useEffect, useState } from "react";
import {
  GetSubscriptionRequest,
  Subscription,
  UpdateSubscriptionRequest,
} from "src/api/subscription_pb";
import { subscriptionService } from "src/Rpc";

type setWatcher = (watcher: string) => void;
type update = () => void;

function useSubscription(
  id: string
): [Subscription | undefined, setWatcher, update] {
  const [subscription, setSubscription] =
    useState<Subscription | undefined>(undefined);
  const [origin, setOrigin] = useState<Subscription | undefined>(undefined);

  useEffect(() => {
    const req = new GetSubscriptionRequest();
    req.setName("subscriptions/" + id);
    subscriptionService.getSubscription(req, {}, (err, res) => {
      if (err === null) {
        setSubscription(res);
        setOrigin(res.clone());
      } else {
        setSubscription(undefined);
        setOrigin(undefined);
      }
    });
  }, [id]);

  return [
    subscription,
    (watcher: string) => {
      if (!subscription) {
        return;
      }
      subscription.setWatcher(watcher);
      setSubscription(subscription);
    },
    () => {
      if (!subscription || !origin) {
        return;
      }
      const mask = new FieldMask();
      if (subscription.getWatcher() !== origin.getWatcher()) {
        mask.addPaths("watcher");
      }
      if (mask.getPathsList().length === 0) {
        return;
      }

      const req = new UpdateSubscriptionRequest();
      req.setUpdateMask(mask);
      req.setName(subscription.getName());
      req.setSubscription(subscription);

      subscriptionService.updateSubscription(req, {}, (err, res) => {
        setSubscription(res);
        setOrigin(res.clone());
      });
    },
  ];
}

export default useSubscription;
