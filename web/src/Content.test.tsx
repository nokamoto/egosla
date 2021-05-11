import React from "react";
import { fireEvent, render } from "@testing-library/react";
import Content from "./Content";
import { watcherService } from "./Rpc";
import {
  CreateWatcherRequest,
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

test("cancel to add a watcher", () => {
  const createWatcher = jest.fn();
  jest.spyOn(watcherService, "createWatcher").mockImplementation(createWatcher);

  const { getByTestId } = render(<Content newChipKeys={[""]} />);

  fireEvent.click(getByTestId("open-addwatch"));
  fireEvent.click(getByTestId("cancel"));

  expect(createWatcher).toHaveBeenCalledTimes(0);
});
