import React, { useEffect, useState, MouseEvent, ChangeEvent } from "react";
import Paper from "@material-ui/core/Paper";
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from "@material-ui/core/styles";
import WatcherDialog from "src/watchers/WatcherDialog";
import { subscriptionService, watcherService } from "src/Rpc";
import {
  CreateWatcherRequest,
  DeleteWatcherRequest,
  ListWatcherRequest,
  UpdateWatcherRequest,
  Watcher,
} from "src/api/watcher_pb";
import { FieldMask } from "google-protobuf/google/protobuf/field_mask_pb";
import WatcherTable from "src/watchers/WatcherTable";
import StandardAppBar from "src/standard/StandardAppBar";
import {
  CreateSubscriptionRequest,
  Subscription,
} from "src/api/subscription_pb";
import { useHistory } from "react-router-dom";
import useStandardMenuList from "src/standard/useStandardMenuList";

const styles = (theme: Theme) =>
  createStyles({
    paper: {
      maxWidth: 936,
      margin: "auto",
      overflow: "hidden",
    },
  });

interface contentProps extends WithStyles<typeof styles> {
  // Keycodes for ChipInput.
  newChipKeys: string[];
}

function WatcherContent(props: contentProps) {
  const { classes } = props;

  const [refresh, setRefresh] = useState(false);
  const [open, setOpen] = useState(false);
  const [updateOpen, setUpdateOpen] = useState(false);
  const [updateKeywords, setUpdateKeywords] = useState<string[]>([]);
  const [updateWatcherName, setUpdateWatcherName] = useState<string>("");
  const [keywords, setKeywords] = useState<string[]>([]);
  const [watchers, setWatchers] = useState<Watcher[]>([]);
  const [anchorEl, openMenu, closeMenu] = useStandardMenuList();
  const [search, setSearch] = useState<string>("");
  const history = useHistory();

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

  const handleClickUpdateMenu = (
    watcherName: string,
    event: MouseEvent<HTMLElement>
  ) => {
    closeMenu();
    setUpdateOpen(true);
    setUpdateWatcherName(watcherName);

    const found = watchers.filter((w) => w.getName() === watcherName);
    if (found.length !== 1) {
      return;
    }
    setUpdateKeywords(found[0].getKeywordsList());
  };

  const deleteWatcher = (watcherName: string, _: MouseEvent<HTMLElement>) => {
    closeMenu();

    const req = new DeleteWatcherRequest();
    req.setName(watcherName);
    watcherService.deleteWatcher(req, {}, (err, res) => {
      setWatchers(watchers.filter((w) => w.getName() !== watcherName));
    });
  };

  const handleReload = () => {
    setRefresh(!refresh);
  };

  const handleSearch = (event: ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
  };

  const handleSubscribe = (watcherName: string, _: MouseEvent<HTMLElement>) => {
    closeMenu();

    const subscription = new Subscription();
    subscription.setWatcher(watcherName);

    const req = new CreateSubscriptionRequest();
    req.setSubscription(subscription);
    subscriptionService.createSubscription(req, {}, (err, res) => {
      history.push("/subscriptions");
    });
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
      <StandardAppBar
        handleClickOpen={handleClickOpen}
        handleReload={handleReload}
        handleSearch={handleSearch}
        searchPlaceholder="Search by name or keywords"
        addText="Add Watcher"
      />
      <WatcherTable
        handleClick={openMenu}
        handleClose={closeMenu}
        handleDelete={deleteWatcher}
        handleUpdate={handleClickUpdateMenu}
        handleSubscribe={handleSubscribe}
        anchorEl={anchorEl}
        watchers={watchers}
        search={search}
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

export default withStyles(styles)(WatcherContent);
