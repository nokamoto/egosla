import React from "react";
import { fireEvent, render, within } from "@testing-library/react";
import WatcherContent from "src/watchers/WatcherContent";
import { subscriptionService, watcherService } from "src/Rpc";
import {
  DeleteWatcherRequest,
  ListWatcherRequest,
  ListWatcherResponse,
  Watcher,
} from "src/api/watcher_pb";
import {
  CreateSubscriptionRequest,
  Subscription,
} from "src/api/subscription_pb";
import { createMemoryHistory } from "history";
import { Router } from "react-router-dom";

test("gets watchers", () => {
  const listWatcher = jest.fn().mockImplementation((x, y, callback) => {
    const watcher = new Watcher();
    watcher.setKeywordsList(["foo", "bar"]);
    const res = new ListWatcherResponse();
    res.addWatchers(watcher);

    callback(null, res);
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const { getByText } = render(<WatcherContent newChipKeys={[]} />);

  const list = new ListWatcherRequest();
  list.setPageSize(100);

  expect(listWatcher).toHaveBeenCalledTimes(1);
  expect(listWatcher.mock.calls[0][0]).toEqual(list);

  expect(getByText("foo")).toBeInTheDocument();
  expect(getByText("bar")).toBeInTheDocument();
});

test("adds a new watcher", () => {
  const history = createMemoryHistory();
  const { getByTestId } = render(
    <Router history={history}>
      <WatcherContent newChipKeys={[]} />
    </Router>
  );

  expect(history.location.pathname).toEqual("/");

  fireEvent.click(getByTestId("open-add"));

  expect(history.location.pathname).toEqual("/watchers/new");
});

test("deletes a watcher", () => {
  const listWatcher = jest.fn().mockImplementation((x, y, callback) => {
    const foo = new Watcher();
    foo.setName("foo");
    foo.setKeywordsList(["bar"]);
    const baz = new Watcher();
    baz.setName("baz");
    baz.setKeywordsList(["qux"]);
    const res = new ListWatcherResponse();
    res.addWatchers(foo);
    res.addWatchers(baz);

    callback(null, res);
  });

  const deleteWatcher = jest.fn().mockImplementation((x, y, callback) => {
    callback(null, null);
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);
  jest.spyOn(watcherService, "deleteWatcher").mockImplementation(deleteWatcher);

  const { queryByText, getByText, getAllByTestId } = render(
    <WatcherContent newChipKeys={[]} />
  );

  expect(getByText("foo")).toBeInTheDocument();
  expect(getByText("baz")).toBeInTheDocument();

  const menus = getAllByTestId("open-menu");
  const del = getAllByTestId("delete");

  expect(menus.length).toEqual(2);
  expect(del.length).toEqual(2);

  fireEvent.click(menus[0]);
  fireEvent.click(del[0]);

  expect(deleteWatcher).toHaveBeenCalledTimes(1);

  const expected = new DeleteWatcherRequest();
  expected.setName("foo");

  expect(deleteWatcher.mock.calls[0][0]).toEqual(expected);

  expect(queryByText("foo")).not.toBeInTheDocument();
  expect(getByText("baz")).toBeInTheDocument();
});

test("updates a watcher", () => {
  const listWatcher = jest.fn().mockImplementation((x, y, callback) => {
    const foo = new Watcher();
    foo.setName("foo");
    foo.setKeywordsList(["bar"]);
    const baz = new Watcher();
    baz.setName("baz");
    baz.setKeywordsList(["qux"]);
    const res = new ListWatcherResponse();
    res.addWatchers(foo);
    res.addWatchers(baz);

    callback(null, res);
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const history = createMemoryHistory();
  const { getAllByTestId } = render(
    <Router history={history}>
      <WatcherContent newChipKeys={[]} />
    </Router>
  );

  expect(history.location.pathname).toEqual("/");

  const menus = getAllByTestId("open-menu");
  const update = getAllByTestId("update");

  expect(menus.length).toEqual(2);
  expect(update.length).toEqual(2);

  fireEvent.click(menus[0]);
  fireEvent.click(update[0]);

  expect(history.location.pathname).toEqual("/foo");
});

test("subscribes a watcher", () => {
  const listWatcher = jest.fn().mockImplementation((x, y, callback) => {
    const foo = new Watcher();
    foo.setName("foo");
    foo.setKeywordsList(["bar"]);
    const baz = new Watcher();
    baz.setName("baz");
    baz.setKeywordsList(["qux"]);
    const res = new ListWatcherResponse();
    res.addWatchers(foo);
    res.addWatchers(baz);

    callback(null, res);
  });

  const createSubscription = jest.fn().mockImplementation((x, y, callback) => {
    callback(null, new Subscription());
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);
  jest
    .spyOn(subscriptionService, "createSubscription")
    .mockImplementation(createSubscription);

  const history = createMemoryHistory();

  const { getAllByTestId } = render(
    <Router history={history}>
      <WatcherContent newChipKeys={[""]} />
    </Router>
  );

  expect(history.location.pathname).toEqual("/");

  const menus = getAllByTestId("open-menu");
  const subscribe = getAllByTestId("subscribe");

  expect(menus.length).toEqual(2);
  expect(subscribe.length).toEqual(2);

  fireEvent.click(menus[0]);
  fireEvent.click(subscribe[0]);

  const subscription = new Subscription();
  subscription.setWatcher("foo");

  const expected = new CreateSubscriptionRequest();
  expected.setSubscription(subscription);

  expect(createSubscription).toHaveBeenCalledTimes(1);
  expect(createSubscription.mock.calls[0][0]).toEqual(expected);

  expect(history.location.pathname).toEqual("/subscriptions");
});

test("reloads a list of watchers", () => {
  const listWatcher = jest.fn();

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const { getByTestId } = render(<WatcherContent newChipKeys={[]} />);

  expect(listWatcher).toHaveBeenCalledTimes(1);

  fireEvent.click(getByTestId("reload"));

  expect(listWatcher).toHaveBeenCalledTimes(2);
});

test("search watchers", () => {
  const listWatcher = jest.fn().mockImplementation((x, y, callback) => {
    const w1 = new Watcher();
    w1.setName("foo");
    w1.setKeywordsList(["bar"]);
    const w2 = new Watcher();
    w2.setName("baz");
    w2.setKeywordsList(["qux"]);
    const res = new ListWatcherResponse();
    res.addWatchers(w1);
    res.addWatchers(w2);

    callback(null, res);
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const { getByTestId } = render(<WatcherContent newChipKeys={[]} />);

  const search = getByTestId("search");
  fireEvent.input(search, { target: { value: "foo" } });

  const table = getByTestId("watchers-table");
  expect(within(table).getByText("foo")).toBeInTheDocument();
  expect(within(table).queryByText("baz")).not.toBeInTheDocument();
});
