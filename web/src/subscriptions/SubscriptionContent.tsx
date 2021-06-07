import React, { ChangeEvent, useEffect, useState, MouseEvent } from "react";
import Paper from "@material-ui/core/Paper";
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from "@material-ui/core/styles";
import StandardAppBar from "src/standard/StandardAppBar";
import { ListSubscriptionRequest, Subscription } from "src/api/subscription_pb";
import { subscriptionService } from "src/Rpc";
import StandardTable from "src/standard/StandardTable";
import { TableCell, TableRow } from "@material-ui/core";
import StandardMenu from "src/standard/StandardMenu";
import DeleteIcon from "@material-ui/icons/Delete";

const styles = (theme: Theme) =>
  createStyles({
    paper: {
      maxWidth: 936,
      margin: "auto",
      overflow: "hidden",
    },
  });

interface contentProps extends WithStyles<typeof styles> {}

function SubscriptionContent(props: contentProps) {
  const { classes } = props;

  const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);
  const [search, setSearch] = useState<string>("");
  const [refresh, setRefresh] = useState<boolean>(false);
  const [anchorEl, setAnchorEl] = useState<HTMLElement[]>([]);

  const handleSearch = (event: ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
  };

  const visibleSubscriptions = subscriptions.filter(
    (s) => s.getName().includes(search) || s.getWatcher().includes(search)
  );

  const handleReload = () => {
    setRefresh(!refresh);
  };

  const handleClickMenu = (index: number, event: MouseEvent<HTMLElement>) => {
    var els: HTMLElement[] = [];
    els[index] = event.currentTarget;
    setAnchorEl(els);
  };

  const handleCloseMenu = () => {
    setAnchorEl([]);
  };

  useEffect(() => {
    const req = new ListSubscriptionRequest();
    req.setPageSize(100);
    subscriptionService.listSubscription(req, {}, (err, res) => {
      setSubscriptions(res.getSubscriptionsList());
    });
  }, [refresh]);

  return (
    <Paper className={classes.paper}>
      <StandardAppBar
        handleReload={handleReload}
        handleSearch={handleSearch}
        searchPlaceholder="Search by name or watcher name"
      />
      <StandardTable
        length={subscriptions.length}
        visibleLength={visibleSubscriptions.length}
        emptyTypography="No subscriptions for this workspace yet"
        tableHeadRow={
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell align="right">Watcher</TableCell>
            <TableCell></TableCell>
          </TableRow>
        }
        tableRows={visibleSubscriptions.map((subscription, index) => {
          return (
            <TableRow key={index.toString()}>
              <TableCell component="th" scope="row">
                {subscription.getName()}
              </TableCell>
              <TableCell align="right">{subscription.getWatcher()}</TableCell>
              <TableCell align="right">
                <StandardMenu
                  index={index}
                  anchorEl={anchorEl}
                  name={subscription.getName()}
                  handleClick={handleClickMenu}
                  handleClose={handleCloseMenu}
                  items={[
                    {
                      icon: <DeleteIcon fontSize="small" />,
                      dataTestID: "delete",
                      itemText: "Delete",
                      onClick: (
                        name: string,
                        event: MouseEvent<HTMLElement>
                      ) => {
                        setAnchorEl([]);
                      },
                    },
                  ]}
                />
              </TableCell>
            </TableRow>
          );
        })}
      />
    </Paper>
  );
}

export default withStyles(styles)(SubscriptionContent);
