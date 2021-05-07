import React from 'react';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogContent from '@material-ui/core/DialogContent';
import DialogActions from '@material-ui/core/DialogActions';
import Button from '@material-ui/core/Button';
import ChipInput from 'material-ui-chip-input';

interface AddWatcherDialogProps{
    open: boolean
    handleCancel: () => void
    handleWatch: () => void
    setKeywords: (keywords: Array<string>) => void
}

function AddWatcherDialog({open, handleCancel, handleWatch, setKeywords}: AddWatcherDialogProps) {
    return (
        <Dialog open={open} onClose={handleCancel} aria-labelledby="form-dialog-title">
                <DialogTitle id="form-dialog-title">Watch Keywords</DialogTitle>
                <DialogContent>
                  <ChipInput
                    label="Keywords"
                    defaultValue={[]}
                    onChange={setKeywords}
                    newChipKeys={[]}
                    fullWidth={true}
                  />
                </DialogContent>
                <DialogActions>
                  <Button onClick={handleCancel} color="primary">
                    Cancel
                  </Button>
                  <Button onClick={handleWatch} color="primary">
                    Watch :eyes:
                  </Button>
                </DialogActions>
              </Dialog>
    );
}

export default AddWatcherDialog;
