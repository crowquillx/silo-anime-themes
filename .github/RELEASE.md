AnimeThemes v0.1.12

Fixes v0.1.11, which exited at startup with `plugin config schema "settings" has invalid json_schema`. The Silo plugin SDK requires each setting to declare a single JSON type, so the settings that accept either text or JSON now declare their two forms with `anyOf`. Update from v0.1.11 to this release.

As in v0.1.11, the list settings take IDs separated by commas or spaces, such as `1, 3`, or a JSON array. **Writable roots and path mappings** and **Provider settings** take a JSON object.

Linux amd64 and arm64 binaries and checksums are attached.
