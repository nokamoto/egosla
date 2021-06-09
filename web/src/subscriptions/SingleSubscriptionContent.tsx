import React from "react";
import Paper from "@material-ui/core/Paper";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import contentStyles from "src/standard/contentStyles";
import { useParams } from "react-router-dom";
import useSubscription from "./useSubscription";
import { TextField } from "@material-ui/core";

interface contentProps extends WithStyles<typeof contentStyles> {}

function SingleSubscriptionContent(props: contentProps) {
  const { classes } = props;

  const { id } = useParams<{ id: string }>();
  const [subscription] = useSubscription(id);

  return (
    <Paper className={classes.paper}>
      {subscription && (
        <div>
          <div className={classes.page}>
            <TextField
              className={classes.textField}
              label="Name"
              defaultValue={subscription.getName()}
              InputProps={{
                readOnly: true,
              }}
              variant="outlined"
            />
          </div>
          <div className={classes.page}>
            <TextField
              className={classes.textField}
              label="Watcher"
              defaultValue={subscription.getWatcher()}
              InputProps={{
                readOnly: true,
              }}
              variant="outlined"
            />
          </div>
        </div>
      )}
    </Paper>
  );
}

export default withStyles(contentStyles)(SingleSubscriptionContent);
