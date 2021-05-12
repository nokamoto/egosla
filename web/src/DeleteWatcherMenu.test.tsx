import React from "react";
import { fireEvent, render } from "@testing-library/react";
import DeleteWatcherMenu from "./DeleteWatcherMenu";

test("deletes a watcher", () => {
  var anchorEl: HTMLElement[] = [];

  const handleClick = jest
    .fn()
    .mockImplementation((index, e) => (anchorEl[index] = e));
  const handleClose = jest.fn();
  const handleDelete = jest.fn();

  const { getByTestId } = render(
    <DeleteWatcherMenu
      index={10}
      anchorEl={anchorEl}
      watcherName="foo"
      handleClick={handleClick}
      handleClose={handleClose}
      handleDelete={handleDelete}
    />
  );

  fireEvent.click(getByTestId("open-menu"));
  fireEvent.click(getByTestId("delete"));

  expect(handleClick).toHaveBeenCalledTimes(1);
  expect(handleClick.mock.calls[0][0]).toEqual(10);

  expect(handleDelete).toHaveBeenCalledTimes(1);
  expect(handleDelete.mock.calls[0][0]).toEqual("foo");

  expect(handleClose).toHaveBeenCalledTimes(0);
});
