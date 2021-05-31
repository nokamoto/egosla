import React from "react";
import { fireEvent, render, within } from "@testing-library/react";
import WatcherContent from "src/watchers/WatcherContent";
import { subscriptionService, watcherService } from "src/Rpc";
import {
  CreateWatcherRequest,
  DeleteWatcherRequest,
  ListWatcherRequest,
  ListWatcherResponse,
  UpdateWatcherRequest,
  Watcher,
} from "src/api/watcher_pb";
import { FieldMask } from "google-protobuf/google/protobuf/field_mask_pb";
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
  const watcher = new Watcher();
  watcher.setKeywordsList(["foo", "bar"]);

  const createWatcher = jest.fn().mockImplementation((x, y, callback) => {
    callback(null, watcher);
  });

  const listWatcher = jest.fn();
  jest.spyOn(watcherService, "createWatcher").mockImplementation(createWatcher);
  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const { getByTestId } = render(<WatcherContent newChipKeys={["Enter"]} />);

  fireEvent.click(getByTestId("open-add"));

  const keywords = getByTestId("keywords");
  fireEvent.input(keywords, { target: { value: "foo" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.input(keywords, { target: { value: "bar" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.click(getByTestId("watch"));

  const expected = new CreateWatcherRequest();
  expected.setWatcher(watcher);

  expect(createWatcher).toHaveBeenCalledTimes(1);
  expect(createWatcher.mock.calls[0][0]).toEqual(expected);

  const table = getByTestId("watchers-table");
  expect(within(table).getByText("foo")).toBeInTheDocument();
  expect(within(table).getByText("bar")).toBeInTheDocument();

  const list = new ListWatcherRequest();
  list.setPageSize(100);

  expect(listWatcher).toHaveBeenCalledTimes(1);
  expect(listWatcher.mock.calls[0][0]).toEqual(list);
});

test("cancels to add a watcher", () => {
  const createWatcher = jest.fn();
  jest.spyOn(watcherService, "createWatcher").mockImplementation(createWatcher);

  const { getByTestId } = render(<WatcherContent newChipKeys={[""]} />);

  fireEvent.click(getByTestId("open-add"));
  fireEvent.click(getByTestId("cancel"));

  expect(createWatcher).toHaveBeenCalledTimes(0);
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

  const watcher = new Watcher();
  watcher.setKeywordsList(["bar", "quux"]);
  const updateMask = new FieldMask();
  updateMask.addPaths("keywords");
  const expected = new UpdateWatcherRequest();
  expected.setName("foo");
  expected.setWatcher(watcher);
  expected.setUpdateMask(updateMask);

  const updateWatcher = jest.fn().mockImplementation((x, y, callback) => {
    const res = new Watcher();
    res.setName(expected.getName());
    res.setKeywordsList(expected.getWatcher()!.getKeywordsList());
    callback(null, res);
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);
  jest.spyOn(watcherService, "updateWatcher").mockImplementation(updateWatcher);

  const { queryByText, getAllByTestId, getByTestId } = render(
    <WatcherContent newChipKeys={["Enter"]} />
  );

  expect(queryByText("quux")).not.toBeInTheDocument();

  const menus = getAllByTestId("open-menu");
  const update = getAllByTestId("update");

  expect(menus.length).toEqual(2);
  expect(update.length).toEqual(2);

  fireEvent.click(menus[0]);
  fireEvent.click(update[0]);

  const keywords = getByTestId("keywords");
  fireEvent.input(keywords, { target: { value: "quux" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.click(getByTestId("watch"));

  expect(updateWatcher).toHaveBeenCalledTimes(1);
  expect(updateWatcher.mock.calls[0][0]).toEqual(expected);

  const table = getByTestId("watchers-table");
  expect(within(table).getByText("quux")).toBeInTheDocument();
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

    callback(null, res);
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const { getByTestId } = render(<WatcherContent newChipKeys={[]} />);

  const search = getByTestId("search");
  fireEvent.input(search, { value: "foo" });

  const table = getByTestId("watchers-table");
  expect(within(table).getByText("foo")).toBeInTheDocument();
  expect(within(table).queryByText("baz")).not.toBeInTheDocument();
});