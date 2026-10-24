AnimeThemes v0.1.9

Takes the shared runtime from Theme Songs v0.1.11.

Series whose Sonarr release was upgraded were refused as "unclassified file or extra". Silo still lists the deleted file. The plugin now ignores listed files that are gone from disk, and it names any file it still cannot classify.

The status page shows progress while a task runs and keeps the last result visible. A summary counts every season, and the item list shows problems and planned downloads first when it reaches its 1,000-row limit. On an 876-series library the summary covers all 2,211 season targets.

The plugin now logs to Silo's container output: each run's start and result, each download, and one line per kind of problem. Set `SILO_THEME_LOG_LEVEL` to change the level.

Fewer requests per run:
- While Preview only is on, Preview and Download share one run.
- Reconcile no longer reads the whole library or looks up themes.
- Before each download is published, the plugin rechecks only that series instead of the whole library. On an 876-series library that check took nine minutes per download.

Linux amd64 and arm64 binaries and checksums are attached.
