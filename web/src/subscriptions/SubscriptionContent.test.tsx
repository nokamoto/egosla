import React from "react";
import { render, within } from "@testing-library/react";
import SubscriptionContent from "./SubscriptionContent";
import { subscriptionService } from "src/Rpc";
import {
  ListSubscriptionRequest,
  ListSubscriptionResponse,
  Subscription,
} from "src/api/subscription_pb";

test("gets subscriptions", () => {
  const subscription = new Subscription();
  subscription.setName("foo");
  subscription.setWatcher("bar");

  const listSubscription = jest.fn().mockImplementation((x, y, callback) => {
    const res = new ListSubscriptionResponse();
    res.setSubscriptionsList([subscription]);
    callback(null, res);
  });

  jest
    .spyOn(subscriptionService, "listSubscription")
    .mockImplementation(listSubscription);

  const { getByTestId } = render(<SubscriptionContent />);

  const expected = new ListSubscriptionRequest();
  expected.setPageSize(100);

  expect(listSubscription).toHaveBeenCalledTimes(1);
  expect(listSubscription.mock.calls[0][0]).toEqual(expected);

  const table = getByTestId("table");
  expect(within(table).getByText(subscription.getName())).toBeInTheDocument();
  expect(
    within(table).getByText(subscription.getWatcher())
  ).toBeInTheDocument();
});
