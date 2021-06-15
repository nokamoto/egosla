import React from "react";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import contentStyles from "src/standard/contentStyles";
import { Paper } from "@material-ui/core";
import ChipInput from "material-ui-chip-input";

interface contentProps extends WithStyles<typeof contentStyles> {}

function NewWatcherContent(props: contentProps) {
  const { classes } = props;
  return (
    <Paper className={classes.paper}>
      <div>
        <div className={classes.page}>
          <ChipInput
            label="Keywords"
            className={classes.textField}
            defaultValue={[]}
            onChange={() => {}}
            newChipKeys={[]}
            InputProps={{
              inputProps: {
                "data-testid": "keywords",
              },
            }}
            variant="outlined"
          />
        </div>
      </div>
    </Paper>
  );
}

export default withStyles(contentStyles)(NewWatcherContent);
