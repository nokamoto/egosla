import React from "react";
import Dialog from "@material-ui/core/Dialog";
import DialogTitle from "@material-ui/core/DialogTitle";
import DialogContent from "@material-ui/core/DialogContent";
import DialogActions from "@material-ui/core/DialogActions";
import Button from "@material-ui/core/Button";
import ChipInput from "material-ui-chip-input";

interface WatcherDialogProps {
  // Dialog open prop.
  open: boolean;
  // Callback fired when Dialog closed.
  handleCancel: () => void;
  // Callback fired when Dialog closed and requested a rpc call.
  handleWatch: () => void;
  // Callback when input keywords changed.
  setKeywords: (keywords: string[]) => void;
  // Keycodes for ChipInput.
  newChipKeys: string[];
  // A button text.
  buttonText: string;
  // Default keywords for ChipInput.
  defaultKeywords: string[];
}

function WatcherDialog({
  open,
  handleCancel,
  handleWatch,
  setKeywords,
  newChipKeys,
  buttonText,
  defaultKeywords,
}: WatcherDialogProps) {
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
          defaultValue={defaultKeywords}
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
          {buttonText}
        </Button>
      </DialogActions>
    </Dialog>
  );
}

export default WatcherDialog;
