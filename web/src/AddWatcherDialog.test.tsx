import React from "react";
import { fireEvent, render } from "@testing-library/react";
import AddWatcherDialog from "./AddWatcherDialog";

test("inputs keywords and watches those", () => {
  const handleCancel = jest.fn();
  const handleWatch = jest.fn();
  const setKeywords = jest.fn();

  const { getByTestId } = render(
    <AddWatcherDialog
      open={true}
      handleCancel={handleCancel}
      handleWatch={handleWatch}
      setKeywords={setKeywords}
      newChipKeys={["Enter"]}
    />
  );

  const keywords = getByTestId("keywords");
  fireEvent.input(keywords, { target: { value: "foo" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.input(keywords, { target: { value: "bar" } });
  fireEvent.keyDown(keywords, { key: "Enter", code: "Enter" });

  fireEvent.click(getByTestId("watch"));

  expect(handleCancel).toHaveBeenCalledTimes(0);
  expect(handleWatch).toHaveBeenCalledTimes(1);
  expect(setKeywords).toHaveBeenCalledTimes(2);
  expect(setKeywords.mock.calls[0][0]).toEqual(["foo"]);
  expect(setKeywords.mock.calls[1][0]).toEqual(["foo", "bar"]);
});

test("cancels", () => {
  const handleCancel = jest.fn();
  const handleWatch = jest.fn();
  const setKeywords = jest.fn();

  const { getByTestId } = render(
    <AddWatcherDialog
      open={true}
      handleCancel={handleCancel}
      handleWatch={handleWatch}
      setKeywords={setKeywords}
      newChipKeys={[]}
    />
  );

  fireEvent.click(getByTestId("cancel"));

  expect(handleCancel).toHaveBeenCalledTimes(1);
  expect(handleWatch).toHaveBeenCalledTimes(0);
  expect(setKeywords).toHaveBeenCalledTimes(0);
});
