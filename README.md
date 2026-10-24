# Silo AnimeThemes

Downloads anime opening and ending themes from AnimeThemes and saves them as MP3 in the matching season's `theme-music` folder. AniBridge and Anime-Lists map Silo seasons to anime entries.

Part of [Crowquillx Silo Plugins](https://github.com/crowquillx/crowquillx-silo-plugins#install). Supports Linux amd64 and arm64.

## Install

1. In **Administration → Plugins → Catalog**, add the [Crowquillx catalog](https://github.com/crowquillx/crowquillx-silo-plugins#install) and install **AnimeThemes**.
2. Restart Silo. Silo adds plugin tasks to **Scheduled Tasks** only when it starts.
3. Create a key in **Administration → API Keys**.
4. In **Installed**, open the plugin's settings, then **Global Configuration**. Paste the key into **Silo administrator API key** and select **Save config**.
5. Run **Preview AnimeThemes** in **Scheduled Tasks**, then open **AnimeThemes** in the admin menu. The page lists your libraries with their IDs.
6. Enter your anime series libraries in **Selected library IDs**, for example `["1"]`, and save. Run Preview again and check the planned downloads.
7. Turn on **Manual library scans**, or enable the plugin's **themes** source in Autoscan with poll delivery and no path rewrites.
8. Turn off **Preview only** and run **Download AnimeThemes**. With manual scans, run **Reconcile AnimeThemes** after the next library scan.

Silo runs each plugin task when it starts. To run them on a schedule instead, edit their triggers in **Scheduled Tasks**. While **Preview only** is on, Preview and Download share one run, and Reconcile only rechecks themes the plugin downloaded.

Tasks keep running after Silo marks them complete. The AnimeThemes status page shows their progress and the last result, with a summary that counts every item.

Silo writes plugin logs to its container output, such as `docker logs silo`, not to its Logs page. Each run logs its start, its result, each download and one line per kind of problem. Set `SILO_THEME_LOG_LEVEL=warn` on the Silo container to log only problems.

Silo writes themes into your media folders, so they must be writable. Silo's Docker Compose file mounts media read-only. Remove `:ro` from the media volume and recreate the container.

If Theme Songs also runs, add the same library IDs to its **Anime-managed library IDs** so it leaves these libraries to AnimeThemes.

### Optional settings

Leave these empty unless your setup differs from Silo's Docker image.

| Setting | Default |
| --- | --- |
| Silo base URL | Silo's internal API address |
| Primary profile ID | The API key owner's primary profile |
| Persistent state directory | `plugin-state/dev.crowquillx.anime-themes` in `SILO_PLUGIN_CACHE_DIR`, which is `/var/lib/silo/plugins` in Docker. Required when that variable is unset. |
| Writable roots and path mappings | The selected libraries' folders |
| ffprobe executable | `/usr/lib/jellyfin-ffmpeg/ffprobe`, then `ffprobe` on PATH |
| Provider `ffmpeg` | `/usr/lib/jellyfin-ffmpeg/ffmpeg`, then `ffmpeg` on PATH |

Use a dedicated series folder with `Season NN` or `Specials` subfolders. Media files may be symlinks, but the plugin refuses other symlinks inside a media folder.

## Build

Requires Go 1.26 or newer. FFmpeg and ffprobe are needed for the audio conversion tests.

```sh
make test
make build
```

The binary is `bin/plugin`.

## Acknowledgments

- [AnimeThemes](https://animethemes.moe/) for theme metadata and audio.
- [AniBridge](https://github.com/anibridge/anibridge-mappings) and [Anime-Lists](https://github.com/Anime-Lists/anime-lists) for anime mappings.
- [FFmpeg](https://ffmpeg.org/) for audio conversion.
- [Silo plugin SDK](https://github.com/Silo-Server/silo-plugin-sdk) for Silo integration.

[MIT license](LICENSE). Dependency licenses are in [THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md).
