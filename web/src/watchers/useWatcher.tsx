import { FieldMask } from "google-protobuf/google/protobuf/field_mask_pb";
import { useEffect, useState } from "react";
import {
  GetWatcherRequest,
  UpdateWatcherRequest,
  Watcher,
} from "src/api/watcher_pb";
import { watcherService } from "src/Rpc";

type watcher = Watcher | undefined;
type setKeywords = (keywords: string[]) => void;
type update = () => void;

function arrayEqual(x: string[], y: string[]): boolean {
  if (x.length !== y.length) {
    return false;
  }
  return x.every((v, i) => v === y[i]);
}

function useWatcher(id: string): [watcher, setKeywords, update] {
  const [watcher, setWatcher] = useState<Watcher | undefined>(undefined);
  const [origin, setOrigin] = useState<Watcher | undefined>(undefined);

  useEffect(() => {
    const req = new GetWatcherRequest();
    req.setName("watchers/" + id);
    watcherService.getWatcher(req, {}, (err, res) => {
      setWatcher(res);
      setOrigin(res.clone());
    });
  }, [id]);

  return [
    watcher,
    (keywords: string[]) => {
      if (!watcher) {
        return;
      }
      watcher.setKeywordsList(keywords);
      setWatcher(watcher);
    },
    () => {
      if (!watcher || !origin) {
        return;
      }

      const mask = new FieldMask();
      if (!arrayEqual(watcher.getKeywordsList(), origin.getKeywordsList())) {
        mask.addPaths("keywords");
      }

      if (mask.getPathsList().length === 0) {
        return;
      }

      const req = new UpdateWatcherRequest();
      req.setName(watcher.getName());
      req.setWatcher(watcher);
      req.setUpdateMask(mask);

      watcherService.updateWatcher(req, {}, (err, res) => {
        setWatcher(res);
        setOrigin(res.clone());
      });
    },
  ];
}

export default useWatcher;
