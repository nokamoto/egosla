import { createStyles, Theme } from "@material-ui/core/styles";

const contentStyles = (theme: Theme) =>
  createStyles({
    paper: {
      maxWidth: 936,
      margin: "auto",
      overflow: "hidden",
    },
  });

export default contentStyles;
