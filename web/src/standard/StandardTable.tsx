import React from "react";
import Typography from "@material-ui/core/Typography";
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from "@material-ui/core/styles";
import Table from "@material-ui/core/Table";
import TableHead from "@material-ui/core/TableHead";
import TableBody from "@material-ui/core/TableBody";

const styles = (theme: Theme) =>
  createStyles({
    contentWrapper: {
      margin: "40px 16px",
    },
    keyword: {
      marginRight: theme.spacing(1),
    },
  });

interface tableProps extends WithStyles<typeof styles> {
  length: number;
  visibleLength: number;
  emptyTypography: string;
  tableHeadRow: React.ReactNode;
  tableRows: React.ReactNode[];
}

function StandardTable(props: tableProps) {
  const {
    classes,
    length,
    visibleLength,
    emptyTypography,
    tableHeadRow,
    tableRows,
  } = props;

  return (
    <div>
      {length === 0 && (
        <div className={classes.contentWrapper}>
          <Typography color="textSecondary" align="center">
            {emptyTypography}
          </Typography>
        </div>
      )}
      {length > 0 && visibleLength === 0 && (
        <div className={classes.contentWrapper}>
          <Typography color="textSecondary" align="center">
            No results matching search
          </Typography>
        </div>
      )}
      {visibleLength > 0 && (
        <Table aria-label="simple table" data-testid="table">
          <TableHead>{tableHeadRow}</TableHead>
          <TableBody>{tableRows}</TableBody>
        </Table>
      )}
    </div>
  );
}

export default withStyles(styles)(StandardTable);
