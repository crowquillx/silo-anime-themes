# AnimeThemes

An unofficial Silo plugin that maps series seasons to AniDB entries and downloads opening and ending audio from AnimeThemes. AniBridge v3 is the primary mapping source. Anime-Lists supplies fallback mappings when the primary source has no usable match. Ambiguous mappings require an explicit selection.

## Install and configure

Add the [Crowquillx plugins catalog](https://raw.githubusercontent.com/crowquillx/crowquillx-silo-plugins/main/repository.json) in Silo Administration, Plugins, Catalog, Repositories. Install **AnimeThemes**. The [releases page](https://github.com/crowquillx/silo-anime-themes/releases) has Linux amd64 and arm64 binaries, checksums, manifests and documentation bundles. Silo's local upload accepts the raw binary. The tarball is a manual deployment bundle.

The plugin needs `ffprobe`, persistent writable state, and writable media mounts visible to the Silo process. It does not need yt-dlp, FFmpeg, the Theme Songs plugin, Renkei, or an AniBridge service.

Use [config.example.json](config.example.json) for the `settings` entry:

1. Supply the Silo HTTP origin, a dedicated administrator's unscoped API key, its primary profile ID, selected library IDs and a persistent state directory. Silo stores the API key as a declared secret.
2. Set `destinations.allowed_roots` to writable plugin-local roots. If server paths differ, provide reversible `path_mappings` with `server` and `local` prefixes. Both plugins must reach the same physical directory for owner locks to coordinate.
3. Keep `preview_only` enabled. Bind and run Preview, then inspect the administrator status page. It shows the mapped source, selected audio, owner destination and any refusal.
4. Enable the plugin's `themes` Autoscan source with poll delivery, no connection credentials and identity path rewrites. Enable Autoscan globally, use a polling interval of at least 60 seconds, and allow two polls for the baseline handshake.
5. Set `preview_only` to false and run Download. Configure a daily task interval in Silo if desired. Restart Silo if its task-binding screen requests it.

Set `manual_refresh` to true only when you intend to run library scans yourself. After scanning, run Reconcile. Neither an emitted scan event nor a downloaded file is reported as discovered until its exact title and owner appear in the native theme set.

Exclude these libraries from the general Theme Songs plugin using that plugin's `anime_library_ids`, or assign individual series through `exclude_items`. Neither plugin replaces another provider's owner marker or audio.

## Select mappings and audio

| Provider field | Meaning |
| --- | --- |
| `mapping_update_hours` | Snapshot refresh interval, default 24 hours. Failed updates retain the last valid snapshot. |
| `op`, `ed` | Include openings and endings. Both default to enabled. |
| `all_distinct` | Include every distinct selected theme, rather than one preferred theme per anime entry. |
| `allow_spoiler`, `allow_nsfw`, `allow_overlap` | Explicit opt-ins. All are false by default. |
| `manual_mappings` | Map `tvdb_show:<id>:s<season>` or `tmdb_show:<id>:s<season>` to an array of AniDB IDs. An empty array disables that mapping. |
| `selection_overrides` | Map `<Silo series ID>:s<season>` to selected `animethemes-<theme ID>-<audio ID>` identifiers. Safety filters still apply. |
| `series_fallback` | Map Silo series IDs to explicitly selected AniDB IDs for the series page. Season downloads do not choose this implicitly. |
| `audio_hosts` | Exact allowed audio hosts, default `a.animethemes.moe`. |
| `anibridge_url`, `anime_lists_url` | Optional HTTPS snapshot URLs. Defaults use the published AniBridge v3 snapshot and Anime-Lists full XML. |

For example, an explicit season mapping can include both cours:

```json
{
  "manual_mappings": {
    "tvdb_show:262954:s2": ["10206", "10835"]
  },
  "op": true,
  "ed": true,
  "all_distinct": true
}
```

Mappings include provenance and a snapshot digest. The JoJo fixtures cover season 2's two anime entries and season 5's three cours. Mapping fixtures do not imply that every corresponding audio asset is currently available. API lookups use exact AniDB resource IDs, pagination, bounded caching and rate-aware retries. Selection deduplicates stable theme/audio identities.

The selected set represents the mapped season, including its cours. It does not adapt to only the downloaded episodes. Silo plays the season's theme set; it does not switch themes at episode-range boundaries.

## Request pacing

AnimeThemes API requests start at most once per second, below its [documented 90 requests per minute](https://github.com/AnimeThemes/animethemes-api-docs/blob/main/docs/jsonapi/intro/ratelimiting/index.md). Audio downloads and mapping snapshot requests use the same one-second minimum per origin, including redirected requests. Silo catalog reads start at most twice per second per plugin process. Cached API results and the daily mapping refresh avoid unnecessary requests.

HTTP requests share per-origin cooldowns across tasks and newly configured clients within the plugin process. `Retry-After` seconds and HTTP dates are honored in full. Exhausted quota headers honor the reset timestamp, including AnimeThemes’ millisecond timestamps. A long cooldown defers work instead of sending an early retry. Requests without a retry header use conservative backoff; attempts and deadlines remain bounded. A failed snapshot update retains the last valid mapping. Restarting the plugin resets in-memory cooldowns.

## Destination rules

Use a dedicated series root with existing conventional `Season NN` folders. Each season destination must contain that season's authoritative episode files and no other series or season. One episode subdirectory below the season is supported. Unknown videos, missing files, unclassified extras, ambiguous copies, symlinks and conflicting observed roots refuse placement. Specials need an existing exclusive `Season 00` directory and an unambiguous mapping.

`single_season_flat_fallback` is disabled by default. If enabled, one populated regular season with no other season or specials can use its validated series root. Preview labels it **series-owned theme, inherited by season/episodes**. The series page plays it too. Every normal run rechecks this proof. A later second season freezes the managed fallback as stale; files are never silently deleted or relocated.

`destination_overrides` uses `library_id/item_id/season/copy_root` keys and server-path values. Overrides still need complete inventory and on-disk checks. They cannot make flat mixed-season media support separate season themes.

Both plugins use protocol 1 `.silo-theme-download.lock` and `.silo-theme-download-owner.json` sidecars in the owner directory. Do not remove them. Use one installation of each plugin, persistent state and reliable local advisory locks. Independent clustered writers and unverified network-filesystem locking are unsupported.

Audio keeps its source format and goes directly into `OWNER/theme-music/`. The writer probes for audio with no video streams, stages on the destination filesystem, records a durable intent, then publishes without replacing an existing file. Manual `theme.*`, unowned audio and edited managed files are preserved. Selection changes leave obsolete files for operator review.

A new destination receives one intended theme first. Further selected themes wait for exact discovery. A failed owner confirmation stops further writes there. After a source reset, let Autoscan establish its new baseline and run Reconcile to request discovery of existing managed files. Back up state with the media; deleting state does not grant ownership of old files.

## Build and test

Requires Go 1.26.0 or later. SDK v0.17.0 and the shared download engine are pinned in `go.mod`. Each binary includes that engine and runs independently.

```sh
go test -race ./...
go vet ./...
make build
bin/plugin manifest
bin/plugin preview /path/to/private-config.json
make build-all
```

The same CLI accepts `sync` and `reconcile`. Keep credential files outside the repository. CI uses controlled mapping and API fixtures. Live provider availability is checked separately. [docs/compatibility.json](docs/compatibility.json) records the tested API contract and owner coordination protocol.

## Sources and license

MIT for this plugin. See [LICENSE](LICENSE) and [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md). Mapping and selection code is an independent implementation; it does not redistribute Renkei code.

Runtime data sources are [AniBridge mappings](https://github.com/anibridge/anibridge-mappings), [Anime-Lists](https://github.com/Anime-Lists/anime-lists), and [AnimeThemes](https://animethemes.moe/). Mapping snapshots and theme audio are fetched at runtime and are not bundled. This plugin is independent of Silo Server and those projects.

See [validation results](docs/validation.md) for live checks and provider availability limits.
