import React from "react";
import Paper from "@material-ui/core/Paper";
import {
  createStyles,
  Theme,
  withStyles,
  WithStyles,
} from "@material-ui/core/styles";

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

  return <Paper className={classes.paper}></Paper>;
}

export default withStyles(styles)(SubscriptionContent);
