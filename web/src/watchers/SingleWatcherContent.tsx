import React from "react";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import contentStyles from "src/standard/contentStyles";
import { Paper, Grid, TextField } from "@material-ui/core";
import { useParams } from "react-router-dom";
import useWatcher from "src/watchers/useWatcher";
import ChipInput from "material-ui-chip-input";

interface contentProps extends WithStyles<typeof contentStyles> {
  // ChipInput.newChipKeys for testing.
  newChipKeys?: string[];
}

function SingleWatcherContent(props: contentProps) {
  const { classes, newChipKeys } = props;
  const { id } = useParams<{ id: string }>();

  const [watcher, setKeywords] = useWatcher(id);

  return (
    <Paper className={classes.paper}>
      {watcher && (
        <div className={classes.page}>
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <TextField
                label="Name"
                defaultValue={watcher.getName()}
                InputProps={{
                  readOnly: true,
                }}
                inputProps={{
                  "data-testid": "name",
                }}
                variant="outlined"
                fullWidth={true}
              />
            </Grid>
            <Grid item xs={6}></Grid>
            <Grid item xs={12}>
              <ChipInput
                label="Keywords"
                defaultValue={watcher.getKeywordsList()}
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
          </Grid>
        </div>
      )}
    </Paper>
  );
}

export default withStyles(contentStyles)(SingleWatcherContent);
