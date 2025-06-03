# Gradual rollout / AB-testing

Demonstration of how to do simple A/B-testing or experimentation, without an
external system. Can be useful for controlling rollout of new features to a
certain number of users, without limiting the number of experiments, and making
sure that every user is getting the same behaviour every time, even as the
rollout scales up.

IDP implemented this first for doing gradual rollout of a new MFA-flow for
logins in TypeScript: <https://github.com/coopnorge/idp-auth0-config/pull/3625>

We then ported this to Go:
<https://github.com/coopnorge/idp-user-service/pull/3059>

This is functionality that we believe should exist in an opinionated library in
Coop, either by calling out to a remote service, or by using this (or similar)
implementations, but it requires a more thorough design to cater to different
audiences.
