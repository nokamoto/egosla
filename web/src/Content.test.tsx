import React from "react";
import { fireEvent, render } from "@testing-library/react";
import Content from "./Content";
import { watcherService } from "./Rpc";
import {
  CreateWatcherRequest,
  DeleteWatcherRequest,
  ListWatcherRequest,
  ListWatcherResponse,
  Watcher,
} from "./api/service_pb";

test("gets watchers", () => {
  const listWatcher = jest.fn().mockImplementation((x, y, callback) => {
    const watcher = new Watcher();
    watcher.setKeywordsList(["foo", "bar"]);
    const res = new ListWatcherResponse();
    res.addWatchers(watcher);

    callback(null, res);
  });

  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const { getByText } = render(<Content newChipKeys={[]} />);

  const list = new ListWatcherRequest();
  list.setPageSize(100);

  expect(listWatcher).toHaveBeenCalledTimes(1);
  expect(listWatcher.mock.calls[0][0]).toEqual(list);

  expect(getByText("foo")).toBeInTheDocument();
  expect(getByText("bar")).toBeInTheDocument();
});

test("adds a new watcher", () => {
  const createWatcher = jest.fn();
  const listWatcher = jest.fn();
  jest.spyOn(watcherService, "createWatcher").mockImplementation(createWatcher);
  jest.spyOn(watcherService, "listWatcher").mockImplementation(listWatcher);

  const { getByTestId, getByText } = render(
    <Content newChipKeys={["Enter"]} />
  );

  fireEvent.click(getByTestId("open-addwatch"));

  const keywords = getByTestId("keywords");
  fireEvent.input(keywords, { target: { value: "foo" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.input(keywords, { target: { value: "bar" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.click(getByTestId("watch"));

  const watcher = new Watcher();
  watcher.setKeywordsList(["foo", "bar"]);
  const expected = new CreateWatcherRequest();
  expected.setWatcher(watcher);

  expect(createWatcher).toHaveBeenCalledTimes(1);
  expect(createWatcher.mock.calls[0][0]).toEqual(expected);

  expect(getByText("foo")).toBeInTheDocument();
  expect(getByText("bar")).toBeInTheDocument();

  const list = new ListWatcherRequest();
  list.setPageSize(100);

  expect(listWatcher).toHaveBeenCalledTimes(1);
  expect(listWatcher.mock.calls[0][0]).toEqual(list);
});

test("cancels to add a watcher", () => {
  const createWatcher = jest.fn();
  jest.spyOn(watcherService, "createWatcher").mockImplementation(createWatcher);

  const { getByTestId } = render(<Content newChipKeys={[""]} />);

  fireEvent.click(getByTestId("open-addwatch"));
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
    <Content newChipKeys={[]} />
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
