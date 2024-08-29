# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Bump fluent-bit from `3.0.6` to `3.1.6`.

## [0.3.0] - 2024-05-28

### Added

- Documentation on how to upgrade fluent-bit.

### Changed

- This PR breaks the existing setup relying on an upstream fork to use only a Dockerfile entry and architect to build the image.

### Removed

- Removes the useless upstream code.

## [0.2.0] - 2023-12-18

### Changed

- Upgrade to fluent-bit 2.2.1

## [0.1.0]

### Changed

- Create custom docker file to add missing aws plugins and auditd.

[Unreleased]: https://github.com/giantswarm/fluent-bit/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/giantswarm/fluent-bit/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/fluent-bit/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/fluent-bit/releases/tag/v0.1.0
