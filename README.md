# MS Redirector

This is a small Go application that handles requests coming to the domain
`www.marlexsystems.org` (with and without `www`). It does so because that
used to be an old domain used for the Spanish Tech Weblog *Marlex Systems*,
but now uses the shorter version, `www.marlex.org`.

Since some of the routes weren't redirecting properly due to several changes
of URLs in the old domain, those routes were being 404'd in Google. By
creating this project, then now those routes can be handled properly.

This is installed in a free tier of [Google App Engine](https://appengine.google.com).
