import React from "react";
import { fireEvent, render } from "@testing-library/react";
import WatcherMenu from "src/watchers/WatcherMenu";

test("deletes a watcher", () => {
  var anchorEl: HTMLElement[] = [];

  const handleClick = jest
    .fn()
    .mockImplementation((index, e) => (anchorEl[index] = e));
  const handleClose = jest.fn();
  const handleDelete = jest.fn();
  const handleUpdate = jest.fn();

  const { getByTestId } = render(
    <WatcherMenu
      index={10}
      anchorEl={anchorEl}
      watcherName="foo"
      handleClick={handleClick}
      handleClose={handleClose}
      handleDelete={handleDelete}
      handleUpdate={handleUpdate}
    />
  );

  fireEvent.click(getByTestId("open-menu"));
  fireEvent.click(getByTestId("delete"));

  expect(handleClick).toHaveBeenCalledTimes(1);
  expect(handleClick.mock.calls[0][0]).toEqual(10);

  expect(handleDelete).toHaveBeenCalledTimes(1);
  expect(handleDelete.mock.calls[0][0]).toEqual("foo");

  expect(handleUpdate).toHaveBeenCalledTimes(0);

  expect(handleClose).toHaveBeenCalledTimes(0);
});

test("updates a watcher", () => {
  var anchorEl: HTMLElement[] = [];

  const handleClick = jest
    .fn()
    .mockImplementation((index, e) => (anchorEl[index] = e));
  const handleClose = jest.fn();
  const handleDelete = jest.fn();
  const handleUpdate = jest.fn();

  const { getByTestId } = render(
    <WatcherMenu
      index={10}
      anchorEl={anchorEl}
      watcherName="foo"
      handleClick={handleClick}
      handleClose={handleClose}
      handleDelete={handleDelete}
      handleUpdate={handleUpdate}
    />
  );

  fireEvent.click(getByTestId("open-menu"));
  fireEvent.click(getByTestId("update"));

  expect(handleClick).toHaveBeenCalledTimes(1);
  expect(handleClick.mock.calls[0][0]).toEqual(10);

  expect(handleDelete).toHaveBeenCalledTimes(0);

  expect(handleUpdate).toHaveBeenCalledTimes(1);
  expect(handleUpdate.mock.calls[0][0]).toEqual("foo");

  expect(handleClose).toHaveBeenCalledTimes(0);
});
