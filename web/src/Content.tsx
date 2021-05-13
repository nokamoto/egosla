import React, { useEffect, useState, MouseEvent } from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
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
import AddWatcherDialog from "./AddWatcherDialog";
import { watcherService } from "./Rpc";
import {
  CreateWatcherRequest,
  DeleteWatcherRequest,
  ListWatcherRequest,
  Watcher,
} from "./api/service_pb";
import Table from "@material-ui/core/Table";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import Chip from "@material-ui/core/Chip";
import WatcherMenu from "./WatcherMenu";

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

  const [open, setOpen] = useState(false);
  const [keywords, setKeywords] = useState<string[]>([]);
  const [watchers, setWatchers] = useState<Watcher[]>([]);
  const [anchorEl, setAnchorEl] = useState<HTMLElement[]>([]);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
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
    console.log("todo");
  };

  const deleteWatcher = (watcherName: string, _: MouseEvent<HTMLElement>) => {
    setAnchorEl([]);

    const req = new DeleteWatcherRequest();
    req.setName(watcherName);
    watcherService.deleteWatcher(req, {}, (err, res) => {
      setWatchers(watchers.filter((w) => w.getName() !== watcherName));
    });
  };

  useEffect(() => {
    const req = new ListWatcherRequest();
    req.setPageSize(100);
    watcherService.listWatcher(req, {}, (err, res) => {
      setWatchers(res.getWatchersList());
    });
  }, []);

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
              <AddWatcherDialog
                open={open}
                handleCancel={handleClose}
                handleWatch={handleWatch}
                setKeywords={setKeywords}
                newChipKeys={props.newChipKeys}
              />
              <Tooltip title="Reload">
                <IconButton>
                  <RefreshIcon className={classes.block} color="inherit" />
                </IconButton>
              </Tooltip>
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
      {watchers.length === 0 && (
        <div className={classes.contentWrapper}>
          <Typography color="textSecondary" align="center">
            No watchers for this workspace yet
          </Typography>
        </div>
      )}
      {watchers.length > 0 && (
        <Table aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell align="right">Keywords</TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {watchers.map((watcher, index) => (
              <TableRow key={index.toString()}>
                <TableCell component="th" scope="row">
                  {watcher.getName()}
                </TableCell>
                <TableCell align="right">
                  {watcher.getKeywordsList().map((keyword, index) => (
                    <Chip
                      key={index.toString()}
                      label={keyword}
                      variant="outlined"
                      className={classes.keyword}
                    />
                  ))}
                </TableCell>
                <TableCell align="right">
                  <WatcherMenu
                    index={index}
                    anchorEl={anchorEl}
                    watcherName={watcher.getName()}
                    handleClick={handleClickDeleteMenu}
                    handleClose={handleCloseDeleteMenu}
                    handleDelete={deleteWatcher}
                    handleUpdate={handleClickUpdateMenu}
                  />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </Paper>
  );
}

export default withStyles(styles)(Content);
