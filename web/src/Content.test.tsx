import React from "react";
import { fireEvent, render } from "@testing-library/react";
import Content from "./Content";
import * as Rpc from "./Rpc";

test("adds a new watcher", () => {
  const createWatch = jest.fn();
  jest.spyOn(Rpc, "createWatch").mockImplementation(createWatch);

  const { getByTestId } = render(<Content newChipKeys={["Enter"]} />);

  fireEvent.click(getByTestId("open-addwatch"));

  const keywords = getByTestId("keywords");
  fireEvent.input(keywords, { target: { value: "foo" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.input(keywords, { target: { value: "bar" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.click(getByTestId("watch"));

  expect(createWatch).toHaveBeenCalledTimes(1);
  expect(createWatch.mock.calls[0][0]).toEqual(["foo", "bar"]);
});

test("cancel to add a watcher", () => {
  const createWatch = jest.fn();
  jest.spyOn(Rpc, "createWatch").mockImplementation(createWatch);

  const { getByTestId } = render(<Content newChipKeys={[""]} />);

  fireEvent.click(getByTestId("open-addwatch"));
  fireEvent.click(getByTestId("cancel"));

  expect(createWatch).toHaveBeenCalledTimes(0);
});
