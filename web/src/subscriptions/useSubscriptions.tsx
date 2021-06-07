import { useEffect, useState } from "react";
import {
  DeleteSubscriptionRequest,
  ListSubscriptionRequest,
  Subscription,
} from "src/api/subscription_pb";
import { subscriptionService } from "src/Rpc";

type subscriptionList = Subscription[];
type visibleSubscriptionList = Subscription[];
type deleteSubscription = (name: string) => void;

function useSubscriptions(
  refresh: boolean,
  search: string
): [subscriptionList, visibleSubscriptionList, deleteSubscription] {
  const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);

  useEffect(() => {
    const req = new ListSubscriptionRequest();
    req.setPageSize(100);
    subscriptionService.listSubscription(req, {}, (err, res) => {
      setSubscriptions(res.getSubscriptionsList());
    });
  }, [refresh]);

  return [
    subscriptions,
    subscriptions.filter(
      (s) => s.getName().includes(search) || s.getWatcher().includes(search)
    ),
    (name: string) => {
      const req = new DeleteSubscriptionRequest();
      req.setName(name);

      subscriptionService.deleteSubscription(req, {}, (err, res) => {
        setSubscriptions(subscriptions.filter((s) => s.getName() !== name));
      });
    },
  ];
}

export default useSubscriptions;
