import { useState, MouseEvent } from "react";

type open = (index: number, event: MouseEvent<HTMLElement>) => void;
type close = () => void;

function useStandardMenuList(): [HTMLElement[], open, close] {
  const [anchorEl, setAnchorEl] = useState<HTMLElement[]>([]);

  return [
    anchorEl,
    (index: number, event: MouseEvent<HTMLElement>) => {
      var els: HTMLElement[] = [];
      els[index] = event.currentTarget;
      setAnchorEl(els);
    },
    () => {
      setAnchorEl([]);
    },
  ];
}

export default useStandardMenuList;
