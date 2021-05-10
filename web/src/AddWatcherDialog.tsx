import React from "react";
import Dialog from "@material-ui/core/Dialog";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogContent from "@material-ui/core/DialogContent";
import DialogActions from "@material-ui/core/DialogActions";
import Button from "@material-ui/core/Button";
import ChipInput from "material-ui-chip-input";

interface AddWatcherDialogProps {
  // Dialog open prop.
  open: boolean;
  // Callback fired when Dialog closed.
  handleCancel: () => void;
  // Callback fired when Dialog closed and requested to add watcher.
  handleWatch: () => void;
  // Callback when input keywords changed.
  setKeywords: (keywords: string[]) => void;
  // Keycodes for ChipInput.
  newChipKeys: string[];
}

function AddWatcherDialog({
  open,
  handleCancel,
  handleWatch,
  setKeywords,
  newChipKeys,
}: AddWatcherDialogProps) {
  return (
    <Dialog
      open={open}
      onClose={handleCancel}
      aria-labelledby="form-dialog-title"
    >
      <DialogTitle id="form-dialog-title">Watch Keywords</DialogTitle>
      <DialogContent>
        <ChipInput
          label="Keywords"
          defaultValue={[]}
          onChange={setKeywords}
          newChipKeys={newChipKeys}
          fullWidth={true}
          InputProps={{
            inputProps: {
              "data-testid": "keywords",
            },
          }}
        />
      </DialogContent>
      <DialogActions>
        <Button onClick={handleCancel} color="primary" data-testid="cancel">
          Cancel
        </Button>
        <Button onClick={handleWatch} color="primary" data-testid="watch">
          Watch :eyes:
        </Button>
      </DialogActions>
    </Dialog>
  );
}

export default AddWatcherDialog;
