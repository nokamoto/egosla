import { useEffect, useState } from "react";
import { GetSubscriptionRequest, Subscription } from "src/api/subscription_pb";
import { subscriptionService } from "src/Rpc";

function useSubscription(id: string): [Subscription | undefined] {
  const [subscription, setSubscription] =
    useState<Subscription | undefined>(undefined);

  useEffect(() => {
    const req = new GetSubscriptionRequest();
    req.setName("subscriptions/" + id);
    subscriptionService.getSubscription(req, {}, (err, res) => {
      if (err === null) {
        setSubscription(res);
      } else {
        setSubscription(undefined);
      }
    });
  }, [id]);

  return [subscription];
}

export default useSubscription;
