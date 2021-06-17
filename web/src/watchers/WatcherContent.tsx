import React, { useEffect, useState, MouseEvent, ChangeEvent } from "react";
import Paper from "@material-ui/core/Paper";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import { subscriptionService, watcherService } from "src/Rpc";
import {
  DeleteWatcherRequest,
  ListWatcherRequest,
  Watcher,
} from "src/api/watcher_pb";
import WatcherTable from "src/watchers/WatcherTable";
import StandardAppBar from "src/standard/StandardAppBar";
import {
  CreateSubscriptionRequest,
  Subscription,
} from "src/api/subscription_pb";
import { useHistory } from "react-router-dom";
import useStandardMenuList from "src/standard/useStandardMenuList";
import contentStyles from "src/standard/contentStyles";

interface contentProps extends WithStyles<typeof contentStyles> {
  // Keycodes for ChipInput.
  newChipKeys: string[];
}

function WatcherContent(props: contentProps) {
  const { classes } = props;

  const [refresh, setRefresh] = useState(false);
  const [watchers, setWatchers] = useState<Watcher[]>([]);
  const [anchorEl, openMenu, closeMenu] = useStandardMenuList();
  const [search, setSearch] = useState<string>("");
  const history = useHistory();

  const handleClickOpen = () => {
    history.push("/watchers/new");
  };

  const handleClickUpdateMenu = (
    watcherName: string,
    event: MouseEvent<HTMLElement>
  ) => {
    history.push(watcherName);
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
    </Paper>
  );
}

export default withStyles(contentStyles)(WatcherContent);
