import React, { ChangeEvent } from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Tooltip from "@material-ui/core/Tooltip";
import IconButton from "@material-ui/core/IconButton";
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from "@material-ui/core/styles";
import SearchIcon from "@material-ui/icons/Search";
import RefreshIcon from "@material-ui/icons/Refresh";

const styles = (theme: Theme) =>
  createStyles({
    paper: {
      maxWidth: 936,
      margin: "auto",
      overflow: "hidden",
    },
    searchBar: {
      borderBottom: "1px solid rgba(0, 0, 0, 0.12)",
    },
    searchInput: {
      fontSize: theme.typography.fontSize,
    },
    block: {
      display: "block",
    },
    addUser: {
      marginRight: theme.spacing(1),
    },
  });

interface standardAppBarProps extends WithStyles<typeof styles> {
  // Callback fired when the add button clicked.
  handleClickOpen: () => void;
  // Callback fired when the reload button clicked.
  handleReload: () => void;
  // Callback fired when the search text changed.
  handleSearch: (event: ChangeEvent<HTMLInputElement>) => void;
  // Placeholder text for the search field.
  searchPlaceholder: string;
  // Text for the add button.
  addText: string;
}

function StandardAppBar(props: standardAppBarProps) {
  const {
    classes,
    handleClickOpen,
    handleReload,
    handleSearch,
    searchPlaceholder,
    addText,
  } = props;

  return (
    <AppBar
      className={classes.searchBar}
      position="static"
      color="default"
      elevation={0}
    >
      <Toolbar>
        <Grid container spacing={2} alignItems="center">
          <Grid item>
            <SearchIcon className={classes.block} color="inherit" />
          </Grid>
          <Grid item xs>
            <TextField
              fullWidth
              placeholder={searchPlaceholder}
              InputProps={{
                disableUnderline: true,
                className: classes.searchInput,
              }}
              onChange={handleSearch}
              data-testid="search"
            />
          </Grid>
          <Grid item>
            <Button
              variant="contained"
              color="primary"
              className={classes.addUser}
              onClick={handleClickOpen}
              data-testid="open-add"
            >
              {addText}
            </Button>
            <Tooltip title="Reload">
              <IconButton onClick={handleReload} data-testid="reload">
                <RefreshIcon className={classes.block} color="inherit" />
              </IconButton>
            </Tooltip>
          </Grid>
        </Grid>
      </Toolbar>
    </AppBar>
  );
}

export default withStyles(styles)(StandardAppBar);
