## About

This repository is used to be able to craft a custom fluent-bit image for the purpose of `enabling the exec plugin` as well as `install the auditd package` in the fluent-bit docker image so we can extract the auditd logs from the machines.

This mostly rely on the upstream dockerfile with a few changes annotated with `Giant Swarm custom` in the Dockerfile present in this repository.

## Upgrade

To upgrade the fluent-bit version, you should go to `./Dockerfile`, bump the value of `ARG RELEASE_VERSION=3.0.6` to the wanted fluent-bit version and create a `Release PR` to trigger the circle CI workflow.
