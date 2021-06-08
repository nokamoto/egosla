import React, { ChangeEvent, useState, MouseEvent } from "react";
import Paper from "@material-ui/core/Paper";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import StandardAppBar from "src/standard/StandardAppBar";
import StandardTable from "src/standard/StandardTable";
import { TableCell, TableRow } from "@material-ui/core";
import StandardMenu from "src/standard/StandardMenu";
import DeleteIcon from "@material-ui/icons/Delete";
import useStandardMenuList from "src/standard/useStandardMenuList";
import useSubscriptions from "./useSubscriptions";
import contentStyles from "src/standard/contentStyles";

interface contentProps extends WithStyles<typeof contentStyles> {}

function SubscriptionContent(props: contentProps) {
  const { classes } = props;

  const [search, setSearch] = useState<string>("");
  const [refresh, setRefresh] = useState<boolean>(false);
  const [anchorEl, openMenu, closeMenu] = useStandardMenuList();
  const [subscriptions, visibleSubscriptions, deleteSubscription] =
    useSubscriptions(refresh, search);

  const handleSearch = (event: ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
  };

  const handleReload = () => {
    setRefresh(!refresh);
  };

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
                  handleClick={openMenu}
                  handleClose={closeMenu}
                  items={[
                    {
                      icon: <DeleteIcon fontSize="small" />,
                      dataTestID: "delete",
                      itemText: "Delete",
                      onClick: (
                        name: string,
                        event: MouseEvent<HTMLElement>
                      ) => {
                        closeMenu();
                        deleteSubscription(name);
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

export default withStyles(contentStyles)(SubscriptionContent);
