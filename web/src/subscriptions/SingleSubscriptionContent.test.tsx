import React from "react";
import { render } from "@testing-library/react";
import { Route, Router } from "react-router-dom";
import SingleSubscriptionContent from "./SingleSubscriptionContent";
import { subscriptionService } from "src/Rpc";
import { GetSubscriptionRequest, Subscription } from "src/api/subscription_pb";
import { createMemoryHistory } from "history";

test("gets a subscription", () => {
  const get = jest.fn().mockImplementation((x, y, callback) => {
    const subscription = new Subscription();
    subscription.setName("subscriptions/foo");
    subscription.setWatcher("watchers/bar");

    callback(null, subscription);
  });

  jest.spyOn(subscriptionService, "getSubscription").mockImplementation(get);

  const history = createMemoryHistory();
  history.push("/foo");

  const { getByTestId } = render(
    <Router history={history}>
      <Route path="/:id">
        <SingleSubscriptionContent />
      </Route>
    </Router>
  );

  const expected = new GetSubscriptionRequest();
  expected.setName("subscriptions/foo");

  expect(get).toHaveBeenCalledTimes(1);
  expect(get.mock.calls[0][0]).toEqual(expected);

  const name = getByTestId("name");
  expect(name).toHaveValue("subscriptions/foo");

  const autocomplete = getByTestId("watcher-autocomplete");
  const watcher = autocomplete.querySelector("input");
  expect(watcher).toHaveValue("watchers/bar");
});
