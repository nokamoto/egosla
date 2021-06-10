import { ChangeEvent, useEffect, useState } from "react";
import { Subscription } from "src/api/subscription_pb";
import { ListWatcherRequest, Watcher } from "src/api/watcher_pb";
import { watcherService } from "src/Rpc";

type isopen = boolean;
type options = Watcher[];
type loading = boolean;
type open = () => void;
type close = () => void;
type inputValue = string;
type setInputValue = (e: ChangeEvent<{}>, s: string) => void;

function useWatcherOptions(
  subscription: Subscription | undefined
): [isopen, options, loading, open, close, inputValue, setInputValue] {
  const [open, setOpen] = useState(false);
  const [options, setOptions] = useState<Watcher[]>([]);
  const loading = open && options.length === 0;
  const [inputValue, setInputValue] = useState("");

  useEffect(() => {
    let active = true;

    if (!loading) {
      return undefined;
    }

    const req = new ListWatcherRequest();
    req.setPageSize(100);
    watcherService.listWatcher(req, {}).then((res) => {
      console.log(res);
      if (active) {
        console.log(active);
        setOptions(res.getWatchersList());
      }
    });

    return () => {
      active = false;
    };
  }, [loading]);

  useEffect(() => {
    if (!open) {
      setOptions([]);
    }
  }, [open]);

  useEffect(() => {
    setInputValue(subscription ? subscription.getWatcher() : "");
  }, [subscription]);

  return [
    open,
    options,
    loading,
    () => {
      setOpen(true);
      setInputValue("");
    },
    () => setOpen(false),
    inputValue,
    (e, s) => setInputValue(s),
  ];
}

export default useWatcherOptions;
