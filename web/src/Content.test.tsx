import React from "react";
import { fireEvent, render } from "@testing-library/react";
import Content from "./Content";
import { watcherService } from "./Rpc";
import { CreateWatcherRequest, Watcher } from "./api/service_pb";

test("adds a new watcher", () => {
  const createWatcher = jest.fn();
  jest.spyOn(watcherService, "createWatcher").mockImplementation(createWatcher);

  const { getByTestId } = render(<Content newChipKeys={["Enter"]} />);

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
});

test("cancel to add a watcher", () => {
  const createWatcher = jest.fn();
  jest.spyOn(watcherService, "createWatcher").mockImplementation(createWatcher);

  const { getByTestId } = render(<Content newChipKeys={[""]} />);

  fireEvent.click(getByTestId("open-addwatch"));
  fireEvent.click(getByTestId("cancel"));

  expect(createWatcher).toHaveBeenCalledTimes(0);
});
