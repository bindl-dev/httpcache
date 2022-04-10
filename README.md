httpcache
=========

[![GoDoc](https://godoc.org/github.com/bindl-dev/httpcache?status.svg)](https://godoc.org/github.com/bindl-dev/httpcache)

> This is a Bindl fork of [gregjones/httpcache](https://github.com/gregjones/httpcache) with minimal dependencies by removing in-tree cache backends, except for diskcache.
> Consider using the upstream project directly should you desire to use other in-tree backends.

Package httpcache provides a http.RoundTripper implementation that works as a mostly [RFC 7234](https://tools.ietf.org/html/rfc7234) compliant cache for http responses.

It is only suitable for use as a 'private' cache (i.e. for a web-browser or an API-client and not for a shared proxy).


Cache Backends
--------------

- The built-in 'memory' cache stores responses in an in-memory map.
- [`sourcegraph.com/sourcegraph/s3cache`](https://sourcegraph.com/github.com/sourcegraph/s3cache) uses Amazon S3 for storage.
- [`github.com/die-net/lrucache`](https://github.com/die-net/lrucache) provides an in-memory cache that will evict least-recently used entries.
- [`github.com/die-net/lrucache/twotier`](https://github.com/die-net/lrucache/tree/master/twotier) allows caches to be combined, for example to use lrucache above with a persistent disk-cache.
- [`github.com/birkelund/boltdbcache`](https://github.com/birkelund/boltdbcache) provides a BoltDB implementation (based on the [bbolt](https://github.com/coreos/bbolt) fork).

If you implement any other backend and wish it to be linked here, please send a PR editing this file.

License
-------

-	[MIT License](LICENSE.txt)
