import React from "react";
import clsx from "clsx";
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from "@material-ui/core/styles";
import Divider from "@material-ui/core/Divider";
import Drawer, { DrawerProps } from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import PeopleIcon from "@material-ui/icons/People";
import { Omit } from "@material-ui/types";
import SubscriptionsIcon from "@material-ui/icons/Subscriptions";
import { Link, useLocation } from "react-router-dom";

const categories = [
  {
    categoryId: "primary",
    children: [
      { id: "Watcher", icon: <PeopleIcon />, to: "/watchers" },
      { id: "Subscription", icon: <SubscriptionsIcon />, to: "/subscriptions" },
    ],
  },
];

const styles = (theme: Theme) =>
  createStyles({
    categoryHeader: {},
    categoryHeaderPrimary: {
      color: theme.palette.common.white,
    },
    item: {
      paddingTop: 1,
      paddingBottom: 1,
      color: "rgba(255, 255, 255, 0.7)",
      "&:hover,&:focus": {
        backgroundColor: "rgba(255, 255, 255, 0.08)",
      },
    },
    itemCategory: {
      backgroundColor: "#232f3e",
      boxShadow: "0 -1px 0 #404854 inset",
      paddingTop: theme.spacing(2),
      paddingBottom: theme.spacing(2),
    },
    firebase: {
      fontSize: 24,
      color: theme.palette.common.white,
    },
    itemActiveItem: {
      color: "#4fc3f7",
    },
    itemPrimary: {
      fontSize: "inherit",
    },
    itemIcon: {
      minWidth: "auto",
      marginRight: theme.spacing(2),
    },
    divider: {
      marginTop: theme.spacing(2),
    },
  });

export interface NavigatorProps
  extends Omit<DrawerProps, "classes">,
    WithStyles<typeof styles> {}

function Navigator(props: NavigatorProps) {
  const { classes, ...other } = props;

  const location = useLocation();

  return (
    <Drawer variant="permanent" {...other}>
      <List disablePadding>
        <ListItem
          className={clsx(classes.firebase, classes.item, classes.itemCategory)}
          key="title"
          component={Link}
          to="/"
        >
          egosla
        </ListItem>
        {categories.map(({ categoryId, children }) => (
          <React.Fragment key={categoryId + "fragment"}>
            <ListItem
              className={classes.categoryHeader}
              key={categoryId}
            ></ListItem>
            {children.map(({ id, icon, to }) => (
              <ListItem
                key={id}
                button
                className={clsx(
                  classes.item,
                  location.pathname === to && classes.itemActiveItem
                )}
                component={Link}
                to={to}
              >
                <ListItemIcon className={classes.itemIcon}>{icon}</ListItemIcon>
                <ListItemText
                  classes={{
                    primary: classes.itemPrimary,
                  }}
                >
                  {id}
                </ListItemText>
              </ListItem>
            ))}
            <Divider className={classes.divider} />
          </React.Fragment>
        ))}
      </List>
    </Drawer>
  );
}

export default withStyles(styles)(Navigator);
