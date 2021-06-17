import { useEffect, useState } from "react";
import { GetWatcherRequest, Watcher } from "src/api/watcher_pb";
import { watcherService } from "src/Rpc";

type watcher = Watcher | undefined;
type setKeywords = (keywords: string[]) => void;

function useWatcher(id: string): [watcher, setKeywords] {
  const [watcher, setWatcher] = useState<Watcher | undefined>(undefined);

  useEffect(() => {
    const req = new GetWatcherRequest();
    req.setName("watchers/" + id);
    watcherService.getWatcher(req, {}, (err, res) => {
      setWatcher(res);
    });
  }, [id]);

  return [
    watcher,
    (keywords: string[]) => {
      watcher?.setKeywordsList(keywords);
      setWatcher(watcher);
    },
  ];
}

export default useWatcher;
