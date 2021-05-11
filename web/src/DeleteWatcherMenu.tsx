import React, { MouseEvent } from "react";
import IconButton from "@material-ui/core/IconButton";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import DeleteIcon from "@material-ui/icons/Delete";
import MenuItem from "@material-ui/core/MenuItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import Menu from "@material-ui/core/Menu";
import ListItemText from "@material-ui/core/ListItemText";

export interface DeleteWatcherMenuProps {
  index: number;
  handleClick: (index: number, event: MouseEvent<HTMLElement>) => void;
  anchorEl: HTMLElement[];
  handleClose: (event: MouseEvent<HTMLElement>) => void;
  watcherName: string;
  handleDelete: (watcherName: string, event: MouseEvent<HTMLElement>) => void;
}

function DeleteWatcherMenu({
  index,
  handleClick,
  anchorEl,
  handleClose,
  watcherName,
  handleDelete,
}: DeleteWatcherMenuProps) {
  return (
    <div>
      <IconButton
        aria-label="more"
        aria-controls="long-menu"
        aria-haspopup="true"
        onClick={(e) => handleClick(index, e)}
      >
        <MoreVertIcon />
      </IconButton>
      <Menu
        id="long-menu"
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
        <MenuItem onClick={(e) => handleDelete(watcherName, e)}>
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
