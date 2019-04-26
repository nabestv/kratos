# naming

## Project Description

Service discovery, service registration related SDK collection

## status quo

Currently, the B-station open source [Discovery] (https://github.com/bilibili/discovery) service registration and discovery SDK is implemented by default.
However, before using it, please confirm that the discovery service deployment is complete and complete the default configuration of the `fixConfig` method in the discovery.go.

## Use

The `Builder` & `Resolver`&`Registry` interface in `naming` can be implemented for service registration and discovery. For example, the Z station also implements zk.