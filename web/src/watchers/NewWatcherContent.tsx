import React from "react";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import contentStyles from "src/standard/contentStyles";
import { Button, Grid, Paper } from "@material-ui/core";
import ChipInput from "material-ui-chip-input";
import useNewWatcher from "src/watchers/useNewWatcher";
import { Watcher } from "src/api/watcher_pb";
import { useHistory } from "react-router-dom";

interface contentProps extends WithStyles<typeof contentStyles> {
  // ChipInput.newChipKeys for testing.
  newChipKeys?: string[];
}

function NewWatcherContent(props: contentProps) {
  const { classes, newChipKeys } = props;
  const [setKeywords, create] = useNewWatcher();
  const history = useHistory();

  return (
    <Paper className={classes.paper}>
      <div className={classes.page}>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <ChipInput
              label="Keywords"
              defaultValue={[]}
              onChange={setKeywords}
              newChipKeys={newChipKeys ? newChipKeys : []}
              InputProps={{
                inputProps: {
                  "data-testid": "keywords",
                },
              }}
              variant="outlined"
              fullWidth={true}
            />
          </Grid>
          <Grid item xs={12}>
            <div className={classes.buttons}>
              <Button
                className={classes.button}
                onClick={() => {
                  history.push("/watchers");
                }}
                color="secondary"
                data-testid="cancel"
              >
                Cancel
              </Button>
              <Button
                className={classes.button}
                onClick={() => {
                  create((res: Watcher) => {
                    history.push("/" + res.getName());
                  });
                }}
                color="primary"
                variant="contained"
                data-testid="create"
              >
                Create
              </Button>
            </div>
          </Grid>
        </Grid>
      </div>
    </Paper>
  );
}

export default withStyles(contentStyles)(NewWatcherContent);
