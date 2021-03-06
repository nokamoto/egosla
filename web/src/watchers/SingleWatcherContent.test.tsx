import React from "react";
import { fireEvent, render } from "@testing-library/react";
import SingleWatcherContent from "./SingleWatcherContent";
import { createMemoryHistory } from "history";
import { Route, Router } from "react-router-dom";
import { watcherService } from "src/Rpc";
import { UpdateWatcherRequest, Watcher } from "src/api/watcher_pb";
import { FieldMask } from "google-protobuf/google/protobuf/field_mask_pb";

describe("watcher found", () => {
  var get = jest.fn();
  var history = createMemoryHistory();
  var element: React.ReactElement;

  beforeEach(() => {
    get = jest.fn().mockImplementation((x, y, cb) => {
      const w = new Watcher();
      w.setName("foo");
      w.setKeywordsList(["bar", "baz"]);
      cb(null, w);
    });
    jest.spyOn(watcherService, "getWatcher").mockImplementation(get);

    history = createMemoryHistory();
    history.push("/foo");

    element = (
      <Router history={history}>
        <Route path="/:id">
          <SingleWatcherContent newChipKeys={["Enter"]} />
        </Route>
      </Router>
    );
  });

  test("gets a watcher", () => {
    const { getByTestId, getAllByRole } = render(element);

    const name = getByTestId("name");
    expect(name).toHaveValue("foo");

    const keywords = getAllByRole("button");
    expect(keywords.length).toBeGreaterThanOrEqual(2);
    expect(keywords[0]).toHaveTextContent("bar");
    expect(keywords[1]).toHaveTextContent("baz");

    expect(get).toHaveBeenCalledTimes(1);
  });

  test("presses back", () => {
    const { getByTestId } = render(element);

    expect(history.location.pathname).toEqual("/foo");

    fireEvent.click(getByTestId("back"));

    expect(history.location.pathname).toEqual("/watchers");
  });

  test("not update if unchanged", () => {
    const update = jest.fn();
    jest.spyOn(watcherService, "updateWatcher").mockImplementation(update);

    const { getByTestId } = render(element);

    fireEvent.click(getByTestId("update"));

    expect(update).toHaveBeenCalledTimes(0);
  });

  test("updates if changed", () => {
    const update = jest.fn();
    jest.spyOn(watcherService, "updateWatcher").mockImplementation(update);

    const { getByTestId } = render(element);

    const keywords = getByTestId("keywords");
    fireEvent.input(keywords, { target: { value: "qux" } });
    fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

    fireEvent.click(getByTestId("update"));

    const w = new Watcher();
    w.setName("foo");
    w.setKeywordsList(["bar", "baz", "qux"]);

    const mask = new FieldMask();
    mask.setPathsList(["keywords"]);

    const expected = new UpdateWatcherRequest();
    expected.setName("foo");
    expected.setWatcher(w);
    expected.setUpdateMask(mask);

    expect(update).toHaveBeenCalledTimes(1);
    expect(update.mock.calls[0][0]).toEqual(expected);
  });
});
