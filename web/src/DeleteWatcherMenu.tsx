import React, { MouseEvent } from "react";
import IconButton from "@material-ui/core/IconButton";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import DeleteIcon from "@material-ui/icons/Delete";
import MenuItem from "@material-ui/core/MenuItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import Menu from "@material-ui/core/Menu";
import ListItemText from "@material-ui/core/ListItemText";

export interface DeleteWatcherMenuProps {
  // Index for anchorEl.
  index: number;
  // HTML element passing to Menu.
  anchorEl: HTMLElement[];
  // Watcher name handled by Menu.
  watcherName: string;
  // Callback fired when Menu opened.
  handleClick: (index: number, event: MouseEvent<HTMLElement>) => void;
  // Callback fired when Menu closed.
  handleClose: (event: MouseEvent<HTMLElement>) => void;
  // Callback fired when Menu closed and requested to delete the watcher.
  handleDelete: (watcherName: string, event: MouseEvent<HTMLElement>) => void;
}

function DeleteWatcherMenu({
  index,
  anchorEl,
  watcherName,
  handleClick,
  handleClose,
  handleDelete,
}: DeleteWatcherMenuProps) {
  return (
    <div>
      <IconButton
        data-testid="open-menu"
        aria-label="more"
        aria-controls="long-menu"
        aria-haspopup="true"
        onClick={(e) => handleClick(index, e)}
      >
        <MoreVertIcon />
      </IconButton>
      <Menu
        anchorEl={anchorEl[index]}
        keepMounted
        open={Boolean(anchorEl[index])}
        onClose={handleClose}
        PaperProps={{
          style: {
            width: "20ch",
          },
        }}
      >
        <MenuItem
          data-testid="delete"
          onClick={(e) => handleDelete(watcherName, e)}
        >
          <ListItemIcon>
            <DeleteIcon fontSize="small" />
          </ListItemIcon>
          <ListItemText primary="Delete" />
        </MenuItem>
      </Menu>
    </div>
  );
}

export default DeleteWatcherMenu;
