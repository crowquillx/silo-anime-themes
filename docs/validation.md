# Compatibility checks

Validated on 2026-09-24 against the server and SDK commits recorded in [compatibility.json](compatibility.json), using a disposable PostgreSQL database and generated video fixtures. The initial installation and ownership checks were followed by a full YouTube download and synthetic-item playback check for v0.1.4.

- Both raw plugin binaries installed through the native upload API. Configuration, admin status routes, RuntimeHost access and scheduled tasks worked.
- An unscoped administrator API key with explicit primary-profile context read catalog membership and file paths. A scoped library-read key received authorization denials on required routes.
- A direct HTTPS audio fixture passed the downloader and ffprobe checks. The shared writer published it without replacement, and the real scanner assigned it to the expected dedicated movie owner with two versions.
- Separate season 2 and season 3 directories received distinct theme files. Their series root retained a separate fallback. Native and Jellyfin theme sets returned the exact expected owner and title.
- Native playback grants returned audio. Jellyfin GET and POST `PlaybackInfo` returned one authorized audio source for the synthetic theme ID. Following each advertised URL returned HTTP 206 and the requested 100-byte range. The server fix also covers client format constraints, conversion and access denials. A mono-only profile refused incompatible stereo audio. An AAC-only conversion profile received a progressive MP4 stream that ffprobe confirmed contains AAC audio. A full FFmpeg decode passed after bounding audio-only MP4 fragments; a 107-second generated-audio regression catches the previously corrupt output. Anonymous playback negotiation returned 401.
- Autoscan established its marker handshake and discovered a new theme downloaded by the hosted scheduled task. Reconciliation confirmed its exact owner. A repeat run did not download it again.
- Focused server tests for the theme repository, scanner, native theme routes and Jellyfin theme routes passed. The synthetic-item playback fix lives in the server’s Jellyfin compatibility handler. The SDK remains v0.17.0.

The Go race suite and vet cover inventory joins, unsafe layouts and mounts, locking, protected files, publication recovery, scan marker retries and resets, provider failures and mapping ambiguity. Provider HTTP and extractor tests use controlled fixtures.

The live ThemerrDB lookup for TMDB series 1399 resolved an exact YouTube URL. The plugin downloaded and validated a 2,574,764-byte MP3 with yt-dlp 2026.08.19, its EJS package and Node 22.23.2. An end-to-end run through the plugin CLI published the audio, the scanner discovered the exact owner, native and Jellyfin playback served it, and a repeated sync downloaded zero files. The previous failure came from treating yt-dlp’s normal `--max-downloads 1` exit code 101 as an error. Regression tests accept that code only in extraction, require a validated audio file, and still reject missing output and actual failures. No source fallback was used.

Rate-limit regressions cover concurrent pacing, newly configured clients, full Retry-After seconds and dates, exhausted quotas, millisecond reset precision, redirected requests, cancellation and deferred work. Tests also verify a shared YouTube search/download queue, cooldown after throttling, explicit runtime arguments and redacted diagnostics. HTTP pacing covers ThemerrDB, direct audio, Silo, AnimeThemes and both mapping sources.

Live mapping snapshots were reachable. The AnimeThemes API/audio service was unavailable during validation, so a live AnimeThemes download could not be verified. Its mapping, selection and failure behavior passed fixture tests. The live direct-audio test verifies the shared downloader and writer, not AnimeThemes asset availability.

The experimental TelevisionTunes source is disabled. Its TLS certificate failed the live probe; TLS verification is never bypassed.

## MP3 output in v0.1.5

The shared downloader passes real FFmpeg conversion and full decode checks for Ogg/Vorbis, Opus, FLAC, WAV, AAC, M4A and M4B sources, including mono and multichannel inputs. MP3 sources remain byte-identical. MP3 packets in a WAV container are copied without re-encoding. The AnimeThemes adapter also passes a generated Ogg-to-MP3 download test. CI and release jobs require these codec tools.

Regression tests cover missing conversion tools before a download, failed and oversized conversion, invalid output, cancellation with child-process cleanup, and publication with the converted extension. The existing publication/discovery test starts with an Ogg candidate, publishes an MP3, confirms discovery and avoids another download on repeated runs. Previously managed files keep their format and source identity.

## Hosted tasks and setup defaults in v0.1.8

This release takes the shared runtime from Theme Songs v0.1.10; its validation report has the details. On 2026-10-24, Theme Songs v0.1.7 failed every hosted Preview on a 37-movie library because the server gives each scheduled-task `Run` call a 10-second deadline. The shared runtime now runs tasks in the background and returns a `running` report before that deadline. It accepts symlinked catalog media files, defaults the Silo URL, profile, state directory, writable roots and FFmpeg paths, and lists the available libraries when none are selected. On an 876-series anime library, a hosted Preview with only an API key and library ID took nine minutes. `Specials` folders now count as season 0; before that change they blocked 70 series. The remaining refusals on that library came from its read-only media mount. AnimeThemes audio comes from direct downloads, so it does not use the managed yt-dlp and Deno tools.

## Shared runtime from Theme Songs v0.1.11

On 2026-10-24, a hosted Preview of the 876-series library showed live progress, and its summary counted all 2,211 season targets while the item list kept 1,000 rows. Theme Songs' validation report covers deleted upgrade files, logging and single-item revalidation.
