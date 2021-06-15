import React from "react";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import contentStyles from "src/standard/contentStyles";
import { Button, Grid, Paper } from "@material-ui/core";
import ChipInput from "material-ui-chip-input";
import { isClassExpression } from "typescript";

interface contentProps extends WithStyles<typeof contentStyles> {}

function NewWatcherContent(props: contentProps) {
  const { classes } = props;
  return (
    <Paper className={classes.paper}>
      <div className={classes.page}>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <ChipInput
              label="Keywords"
              defaultValue={[]}
              onChange={() => {}}
              newChipKeys={[]}
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
            <div
              style={{
                display: "flex",
                justifyContent: "flex-end",
              }}
            >
              <Button onClick={() => {}} color="secondary" data-testid="cancel">
                Cancel
              </Button>
              <Button
                onClick={() => {}}
                color="primary"
                variant="contained"
                data-testid="watch"
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
