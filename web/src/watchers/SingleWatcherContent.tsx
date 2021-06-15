import React from "react";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import contentStyles from "src/standard/contentStyles";
import { Paper } from "@material-ui/core";
import { useParams } from "react-router-dom";

interface contentProps extends WithStyles<typeof contentStyles> {}

function SingleWatcherContent(props: contentProps) {
  const { classes } = props;
  const { id } = useParams<{ id: string }>();

  return <Paper className={classes.paper}>id={id}</Paper>;
}

export default withStyles(contentStyles)(SingleWatcherContent);
