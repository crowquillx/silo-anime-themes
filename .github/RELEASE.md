AnimeThemes v0.1.8

Fixes tasks that failed on real libraries with "Silo read rate limit exceeds remaining request budget". Silo allows each scheduled task call 10 seconds, and reading a library takes longer. Preview, Download and Reconcile now continue in the background after that call returns; the status page shows the result. A second trigger joins the task already running.

Setup needs only an API key. The Silo URL, profile, state directory and writable roots now have defaults. FFmpeg and ffprobe default to Silo's Docker paths under `/usr/lib/jellyfin-ffmpeg`, falling back to PATH. Without library IDs, Preview lists your libraries with their IDs. The README install steps now cover the server restart Silo needs before plugin tasks appear and the read-only media mount in Silo's Docker Compose file.

Series with a `Specials` folder, the Sonarr and Jellyfin name for season 0, are no longer refused as an unsupported layout. On a test library this blocked 70 of 876 series. Library reads are four times faster; a preview of 876 series takes about nine minutes instead of 45.

Series whose episode files are symlinks, as in debrid and rclone libraries, are no longer refused. Other symlinks are still refused, and the message names the file.

Linux amd64 and arm64 binaries and checksums are attached.
