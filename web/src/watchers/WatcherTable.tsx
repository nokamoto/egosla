import React, { MouseEvent } from "react";
import Typography from "@material-ui/core/Typography";
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from "@material-ui/core/styles";
import { Watcher } from "src/api/watcher_pb";
import Table from "@material-ui/core/Table";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TableBody from "@material-ui/core/TableBody";
import Chip from "@material-ui/core/Chip";
import WatcherMenu from "src/watchers/WatcherMenu";

const styles = (theme: Theme) =>
  createStyles({
    contentWrapper: {
      margin: "40px 16px",
    },
    keyword: {
      marginRight: theme.spacing(1),
    },
  });

export interface WatcherTableProps extends WithStyles<typeof styles> {
  // Watchers retrieved from the backend.
  watchers: Watcher[];
  // A search text string.
  search: string;
  // HTML element passing to Menu.
  anchorEl: HTMLElement[];
  // Callback fired when Menu opened.
  handleClick: (index: number, event: MouseEvent<HTMLElement>) => void;
  // Callback fired when Menu closed.
  handleClose: (event: MouseEvent<HTMLElement>) => void;
  // Callback fired when Menu closed and requested to delete the watcher.
  handleDelete: (watcherName: string, event: MouseEvent<HTMLElement>) => void;
  // Callback fired when Menu closed and requested to update the watcher.
  handleUpdate: (watcherName: string, event: MouseEvent<HTMLElement>) => void;
  // Callback fired when Menu closed and requested to subscribe the watcher.
  handleSubscribe: (
    watcherName: string,
    event: MouseEvent<HTMLElement>
  ) => void;
}

function WatcherTable(props: WatcherTableProps) {
  const {
    classes,
    watchers,
    search,
    anchorEl,
    handleClick,
    handleClose,
    handleDelete,
    handleUpdate,
    handleSubscribe,
  } = props;

  const visibleWatchers = watchers.filter(
    (w) =>
      w.getName().includes(search) ||
      w.getKeywordsList().some((k) => k.includes(search))
  );

  return (
    <div>
      {watchers.length === 0 && (
        <div className={classes.contentWrapper}>
          <Typography color="textSecondary" align="center">
            No watchers for this workspace yet
          </Typography>
        </div>
      )}
      {watchers.length > 0 && visibleWatchers.length === 0 && (
        <div className={classes.contentWrapper}>
          <Typography color="textSecondary" align="center">
            No results matching search
          </Typography>
        </div>
      )}
      {visibleWatchers.length > 0 && (
        <Table aria-label="simple table" data-testid="watchers-table">
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell align="right">Keywords</TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {visibleWatchers.map((watcher, index) => (
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
                    handleClick={handleClick}
                    handleClose={handleClose}
                    handleDelete={handleDelete}
                    handleUpdate={handleUpdate}
                    handleSubscribe={handleSubscribe}
                  />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </div>
  );
}

export default withStyles(styles)(WatcherTable);
