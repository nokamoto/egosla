import React from "react";
import { fireEvent, render } from "@testing-library/react";
import NewWatcherContent from "src/watchers/NewWatcherContent";
import { createMemoryHistory } from "history";
import { Router } from "react-router-dom";
import { watcherService } from "src/Rpc";
import { CreateWatcherRequest, Watcher } from "src/api/watcher_pb";

test("creates a watcher", () => {
  const create = jest.fn().mockImplementation((x, y, cb) => {
    const res = new Watcher();
    res.setName("watchers/foo");
    cb(null, res);
  });

  jest.spyOn(watcherService, "createWatcher").mockImplementation(create);

  const history = createMemoryHistory();

  const { getByTestId } = render(
    <Router history={history}>
      <NewWatcherContent newChipKeys={["Enter"]} />
    </Router>
  );

  expect(history.location.pathname).toEqual("/");

  const keywords = getByTestId("keywords");
  fireEvent.input(keywords, { target: { value: "foo" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.click(getByTestId("create"));

  const created = new Watcher();
  created.setKeywordsList(["foo"]);
  const expected = new CreateWatcherRequest();
  expected.setWatcher(created);

  expect(create).toHaveBeenCalledTimes(1);
  expect(create.mock.calls[0][0]).toEqual(expected);

  expect(history.location.pathname).toEqual("/watchers/foo");
});

test("cancels to create a watcher", () => {
  const history = createMemoryHistory();

  const { getByTestId } = render(
    <Router history={history}>
      <NewWatcherContent />
    </Router>
  );

  expect(history.location.pathname).toEqual("/");

  fireEvent.click(getByTestId("cancel"));

  expect(history.location.pathname).toEqual("/watchers");
});
