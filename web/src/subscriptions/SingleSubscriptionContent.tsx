import React from "react";
import Paper from "@material-ui/core/Paper";
import { withStyles, WithStyles } from "@material-ui/core/styles";
import contentStyles from "src/standard/contentStyles";
import { useParams } from "react-router-dom";
import useSubscription from "./useSubscription";
import { TextField } from "@material-ui/core";
import Autocomplete from "@material-ui/lab/Autocomplete";
import useWatcherOptions from "./useWatcherOptions";
import CircularProgress from "@material-ui/core/CircularProgress";

interface contentProps extends WithStyles<typeof contentStyles> {}

function SingleSubscriptionContent(props: contentProps) {
  const { classes } = props;

  const { id } = useParams<{ id: string }>();
  const [subscription] = useSubscription(id);

  const [isopen, options, loading, open, close, inputValue, setInputValue] =
    useWatcherOptions(subscription);

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
              inputProps={{
                "data-testid": "name",
              }}
              variant="outlined"
            />
          </div>
          <div className={classes.page}>
            <Autocomplete
              style={{ width: 300 }}
              open={isopen}
              onOpen={open}
              onClose={close}
              inputValue={inputValue}
              onInputChange={setInputValue}
              getOptionSelected={(option, value) =>
                option.getName() === value.getName()
              }
              getOptionLabel={(option) => option.getName()}
              options={options}
              loading={loading}
              data-testid="watcher-autocomplete"
              renderInput={(params) => (
                <TextField
                  {...params}
                  className={classes.textField}
                  label="Watcher"
                  variant="outlined"
                  InputProps={{
                    ...params.InputProps,
                    endAdornment: (
                      <React.Fragment>
                        {loading ? (
                          <CircularProgress color="inherit" size={20} />
                        ) : null}
                        {params.InputProps.endAdornment}
                      </React.Fragment>
                    ),
                  }}
                />
              )}
            />
          </div>
        </div>
      )}
    </Paper>
  );
}

export default withStyles(contentStyles)(SingleSubscriptionContent);
