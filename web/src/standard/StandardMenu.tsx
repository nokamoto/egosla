import React, { MouseEvent } from "react";
import IconButton from "@material-ui/core/IconButton";
import MoreVertIcon from "@material-ui/icons/MoreVert";
import MenuItem from "@material-ui/core/MenuItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import Menu from "@material-ui/core/Menu";
import ListItemText from "@material-ui/core/ListItemText";

interface menuItem {
  // Icon for ListItemIcon.
  icon: React.ReactNode;
  // ID for data-testid.
  dataTestID: string;
  // Primary for ListItemText.
  itemText: string;
  // Callback fired when MenuItem selected.
  onClick: (name: string, event: MouseEvent<HTMLElement>) => void;
}

interface menuProps {
  // Index for anchorEl.
  index: number;
  // HTML element passing to Menu.
  anchorEl: HTMLElement[];
  // Resource name handled by Menu.
  name: string;
  // Callback fired when Menu opened.
  handleClick: (index: number, event: MouseEvent<HTMLElement>) => void;
  // Callback fired when Menu closed.
  handleClose: (event: MouseEvent<HTMLElement>) => void;
  // Menu items.
  items: menuItem[];
}

function StandardMenu(props: menuProps) {
  const { handleClick, index, anchorEl, handleClose, items, name } = props;

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
        {items.map((item) => {
          return (
            <MenuItem
              data-testid={item.dataTestID}
              key={item.dataTestID}
              onClick={(e) => item.onClick(name, e)}
            >
              <ListItemIcon>{item.icon}</ListItemIcon>
              <ListItemText primary={item.itemText} />
            </MenuItem>
          );
        })}
      </Menu>
    </div>
  );
}

export default StandardMenu;
