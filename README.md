## Overview
My wife and I use the fabulous [BookBuddy iOS app](https://itunes.apple.com/us/app/bookbuddy-pro/id395149950?mt=8) to catalogue and manage our home library. BookBuddy generates a SQLite 3 database file which is synced between devices via Dropbox. Serenity, written in Go, uses that database to access book information and display it via a RESTful HTTP API and HTML interface. This is useful for a few reasons:
1. Utilizing a mounted Chromebook, we can easily navigate our home library which is organized onto multiple shelves using library-grade labels and the [Library of Congress Classification](https://www.loc.gov/catdir/cpso/lcc.html) classification system.
1. We can share a list of our library's contents with friends and family who may want recommendations or to know if we have a book they're considering purchasing for us.
1. We can maintain and publish book wish lists.

## Quickstart
```
cd frontend && ng build --aot && cd .. && go run server.go
```

## Troubleshooting
Title, author, and ISBN are required by the server. Run the following query to find problematic records:
```
SELECT ZTITLE, ZDISPLAYNAME, ZISBN FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZTITLE LIKE "" OR ZDISPLAYNAME LIKE "" OR ZISBN LIKE "";
```
