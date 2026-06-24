<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/notifications-webpush</h1>
  <p>Web Push (VAPID) channel for togo notifications — push to subscribed browsers.</p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install
```bash
togo install togo-framework/notifications-webpush
```
<!-- /togo-header -->

Set `VAPID_PUBLIC_KEY`, `VAPID_PRIVATE_KEY`, `VAPID_SUBJECT`. Recipients route push tokens as JSON `webpush.Subscription` strings; `PushNotification.ToPush` supplies the title/body.

<!-- togo-sponsors -->
---
<div align="center">
  <h3>Premium sponsors</h3>
  <p><a href="https://id8media.com"><strong>ID8 Media</strong></a> &middot; <a href="https://one-studio.co"><strong>One Studio</strong></a></p>
  <p><sub>Support togo &mdash; <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
