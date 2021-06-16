import { useState } from "react";
import { CreateWatcherRequest, Watcher } from "src/api/watcher_pb";
import { watcherService } from "src/Rpc";

type setKeywords = (keywords: string[]) => void;
type create = (cb: (res: Watcher) => void) => void;

function useNewWatcher(): [setKeywords, create] {
  const [watcher, setWatcher] = useState<Watcher>(new Watcher());
  return [
    (keywords: string[]) => {
      watcher.setKeywordsList(keywords);
      setWatcher(watcher);
    },
    (cb: (res: Watcher) => void) => {
      const req = new CreateWatcherRequest();
      req.setWatcher(watcher);
      watcherService.createWatcher(req, {}, (err, res) => {
        cb(res);
      });
    },
  ];
}

export default useNewWatcher;
