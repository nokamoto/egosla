import { useEffect, useState } from "react";
import { ListWatcherRequest, Watcher } from "src/api/watcher_pb";
import { watcherService } from "src/Rpc";

type isopen = boolean;
type options = Watcher[];
type loading = boolean;
type open = () => void;
type close = () => void;

function useWatcherOptions(): [isopen, options, loading, open, close] {
  const [open, setOpen] = useState(false);
  const [options, setOptions] = useState<Watcher[]>([]);
  const loading = open && options.length === 0;

  useEffect(() => {
    let active = true;

    if (!loading) {
      return undefined;
    }

    const req = new ListWatcherRequest();
    req.setPageSize(100);
    watcherService.listWatcher(req, {}).then((res) => {
      if (active) {
        setOptions(res.getWatchersList());
      }
    });

    return () => {
      active = false;
    };
  });

  useEffect(() => {
    if (!open) {
      setOptions([]);
    }
  }, [open]);

  return [open, options, loading, () => setOpen(true), () => setOpen(false)];
}

export default useWatcherOptions;
