import React from "react";
import { fireEvent, render, within } from "@testing-library/react";
import SubscriptionContent from "./SubscriptionContent";
import { subscriptionService } from "src/Rpc";
import {
  DeleteSubscriptionRequest,
  ListSubscriptionRequest,
  ListSubscriptionResponse,
  Subscription,
} from "src/api/subscription_pb";
import { Empty } from "google-protobuf/google/protobuf/empty_pb";
import { createMemoryHistory } from "history";
import { Router } from "react-router-dom";

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

test("searches subscriptions", () => {
  const listSubscription = jest.fn().mockImplementation((x, y, callback) => {
    const s1 = new Subscription();
    s1.setName("foo");

    const s2 = new Subscription();
    s2.setName("bar");

    const res = new ListSubscriptionResponse();
    res.setSubscriptionsList([s1, s2]);
    callback(null, res);
  });

  jest
    .spyOn(subscriptionService, "listSubscription")
    .mockImplementation(listSubscription);

  const { getByTestId } = render(<SubscriptionContent />);

  const search = getByTestId("search");
  fireEvent.input(search, { target: { value: "foo" } });

  expect(listSubscription).toHaveBeenCalledTimes(1);

  const table = getByTestId("table");
  expect(within(table).getByText("foo")).toBeInTheDocument();
  expect(within(table).queryByText("bar")).not.toBeInTheDocument();
});

test("reload subscriptions", () => {
  const listSubscription = jest.fn();

  jest
    .spyOn(subscriptionService, "listSubscription")
    .mockImplementation(listSubscription);

  const { getByTestId } = render(<SubscriptionContent />);

  fireEvent.click(getByTestId("reload"));

  const expected = new ListSubscriptionRequest();
  expected.setPageSize(100);

  expect(listSubscription).toHaveBeenCalledTimes(2);
  expect(listSubscription.mock.calls[1][0]).toEqual(expected);
});

test("delete subscriptions", () => {
  const listSubscription = jest.fn().mockImplementation((x, y, callback) => {
    const s1 = new Subscription();
    s1.setName("foo");

    const s2 = new Subscription();
    s2.setName("bar");

    const res = new ListSubscriptionResponse();
    res.setSubscriptionsList([s1, s2]);
    callback(null, res);
  });

  const deleteSubscription = jest.fn().mockImplementation((x, y, callback) => {
    callback(null, new Empty());
  });

  jest
    .spyOn(subscriptionService, "listSubscription")
    .mockImplementation(listSubscription);
  jest
    .spyOn(subscriptionService, "deleteSubscription")
    .mockImplementation(deleteSubscription);

  const { getByTestId, getAllByTestId } = render(<SubscriptionContent />);

  expect(listSubscription).toHaveBeenCalledTimes(1);
  expect(deleteSubscription).toHaveBeenCalledTimes(0);

  let table = getByTestId("table");
  expect(within(table).getByText("foo")).toBeInTheDocument();
  expect(within(table).getByText("bar")).toBeInTheDocument();

  const menus = getAllByTestId("open-menu");
  const del = getAllByTestId("delete");

  expect(menus.length).toEqual(2);
  expect(del.length).toEqual(2);

  fireEvent.click(menus[0]);
  fireEvent.click(del[0]);

  const expected = new DeleteSubscriptionRequest();
  expected.setName("foo");

  expect(deleteSubscription).toHaveBeenCalledTimes(1);
  expect(deleteSubscription.mock.calls[0][0]).toEqual(expected);

  table = getByTestId("table");
  expect(within(table).queryByText("foo")).not.toBeInTheDocument();
  expect(within(table).getByText("bar")).toBeInTheDocument();
});

test("updates subscriptions", () => {
  const listSubscription = jest.fn().mockImplementation((x, y, callback) => {
    const s1 = new Subscription();
    s1.setName("foo");

    const s2 = new Subscription();
    s2.setName("bar");

    const res = new ListSubscriptionResponse();
    res.setSubscriptionsList([s1, s2]);
    callback(null, res);
  });

  jest
    .spyOn(subscriptionService, "listSubscription")
    .mockImplementation(listSubscription);

  const history = createMemoryHistory();

  const { getAllByTestId } = render(
    <Router history={history}>
      <SubscriptionContent />
    </Router>
  );

  expect(history.location.pathname).toEqual("/");

  expect(listSubscription).toHaveBeenCalledTimes(1);

  const menus = getAllByTestId("open-menu");
  const update = getAllByTestId("update");

  expect(menus.length).toEqual(2);
  expect(update.length).toEqual(2);

  fireEvent.click(menus[0]);
  fireEvent.click(update[0]);

  expect(history.location.pathname).toEqual("/foo");
});
