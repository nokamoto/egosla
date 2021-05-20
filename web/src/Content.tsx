import React, { useEffect, useState, MouseEvent } from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Paper from "@material-ui/core/Paper";
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
import WatcherDialog from "./WatcherDialog";
import { watcherService } from "./Rpc";
import {
  CreateWatcherRequest,
  DeleteWatcherRequest,
  ListWatcherRequest,
  UpdateWatcherRequest,
  Watcher,
} from "./api/service_pb";
import { FieldMask } from "google-protobuf/google/protobuf/field_mask_pb";
import WatcherTable from "./WatcherTable";

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
    contentWrapper: {
      margin: "40px 16px",
    },
    keyword: {
      marginRight: theme.spacing(1),
    },
  });

export interface ContentProps extends WithStyles<typeof styles> {
  // Keycodes for ChipInput.
  newChipKeys: string[];
}

function Content(props: ContentProps) {
  const { classes } = props;

  const [refresh, setRefresh] = useState(false);
  const [open, setOpen] = useState(false);
  const [updateOpen, setUpdateOpen] = useState(false);
  const [updateKeywords, setUpdateKeywords] = useState<string[]>([]);
  const [updateWatcherName, setUpdateWatcherName] = useState<string>("");
  const [keywords, setKeywords] = useState<string[]>([]);
  const [watchers, setWatchers] = useState<Watcher[]>([]);
  const [anchorEl, setAnchorEl] = useState<HTMLElement[]>([]);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleUpdateClose = () => {
    setUpdateOpen(false);
  };

  const handleUpdate = () => {
    setUpdateOpen(false);

    const watcher = new Watcher();
    watcher.setKeywordsList(updateKeywords);

    const updateMask = new FieldMask();
    updateMask.addPaths("keywords");

    const req = new UpdateWatcherRequest();
    req.setName(updateWatcherName);
    req.setWatcher(watcher);
    req.setUpdateMask(updateMask);

    watcherService.updateWatcher(req, {}, (err, res) => {
      setWatchers(
        watchers.map((v) => {
          if (v.getName() === res.getName()) {
            return res;
          }
          return v;
        })
      );
    });
  };

  const handleWatch = () => {
    setOpen(false);

    const watcher = new Watcher();
    watcher.setKeywordsList(keywords);
    const req = new CreateWatcherRequest();
    req.setWatcher(watcher);

    watcherService.createWatcher(req, {}, (err, res) => {
      setWatchers(watchers.concat(res));
    });
  };

  const handleClickDeleteMenu = (
    index: number,
    event: MouseEvent<HTMLElement>
  ) => {
    var els: HTMLElement[] = [];
    els[index] = event.currentTarget;
    setAnchorEl(els);
  };

  const handleCloseDeleteMenu = () => {
    setAnchorEl([]);
  };

  const handleClickUpdateMenu = (
    watcherName: string,
    event: MouseEvent<HTMLElement>
  ) => {
    setAnchorEl([]);
    setUpdateOpen(true);
    setUpdateWatcherName(watcherName);

    const found = watchers.filter((w) => w.getName() === watcherName);
    if (found.length !== 1) {
      return;
    }
    setUpdateKeywords(found[0].getKeywordsList());
  };

  const deleteWatcher = (watcherName: string, _: MouseEvent<HTMLElement>) => {
    setAnchorEl([]);

    const req = new DeleteWatcherRequest();
    req.setName(watcherName);
    watcherService.deleteWatcher(req, {}, (err, res) => {
      setWatchers(watchers.filter((w) => w.getName() !== watcherName));
    });
  };

  const handleReload = () => {
    setRefresh(!refresh);
  };

  useEffect(() => {
    const req = new ListWatcherRequest();
    req.setPageSize(100);
    watcherService.listWatcher(req, {}, (err, res) => {
      setWatchers(res.getWatchersList());
    });
  }, [refresh]);

  return (
    <Paper className={classes.paper}>
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
                placeholder="Search by name or keywords"
                InputProps={{
                  disableUnderline: true,
                  className: classes.searchInput,
                }}
              />
            </Grid>
            <Grid item>
              <Button
                variant="contained"
                color="primary"
                className={classes.addUser}
                onClick={handleClickOpen}
                data-testid="open-addwatch"
              >
                Add Watcher
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
      <WatcherTable
        handleClick={handleClickDeleteMenu}
        handleClose={handleCloseDeleteMenu}
        handleDelete={deleteWatcher}
        handleUpdate={handleClickUpdateMenu}
        anchorEl={anchorEl}
        watchers={watchers}
      />
      <WatcherDialog
        open={open}
        handleCancel={handleClose}
        handleWatch={handleWatch}
        setKeywords={setKeywords}
        newChipKeys={props.newChipKeys}
        buttonText="Watch :eye:"
        defaultKeywords={[]}
      />
      <WatcherDialog
        open={updateOpen}
        handleCancel={handleUpdateClose}
        handleWatch={handleUpdate}
        setKeywords={setUpdateKeywords}
        newChipKeys={props.newChipKeys}
        buttonText="Update :pen:"
        defaultKeywords={updateKeywords}
      />
    </Paper>
  );
}

export default withStyles(styles)(Content);
