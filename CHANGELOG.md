# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-02-20

### Added
- Initial implementation of dotenv loader
- `Load`, `Overload`, and `Read` functions
- Basic parsing: comments, quotes, export prefix, inline comments
- Basic multi-line value support (double-quoted)
- Error handling with file/line details
- Comprehensive tests covering happy paths, missing file, parse errors, multi-line
- Detailed README with examples, syntax, limitations, and philosophy
- .env.example sample file

### Notes
- Zero dependencies
- Designed for local dev, CLI tools, small services, and AI agents
- No variable expansion or advanced features yet (planned for future)

[0.1.0]: https://github.com/njchilds90/go-dotenv-simple/releases/tag/v0.1.0