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
### Empty Values
Title, author, and ISBN are required by the server. Run the following query to find problematic records:
```
SELECT ZTITLE, ZDISPLAYNAME, ZISBN FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZTITLE LIKE "" OR ZDISPLAYNAME LIKE "" OR ZISBN LIKE "";
```

### Poorly Formatted Author Names
```
SELECT ZTITLE, ZDISPLAYNAME FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZDISPLAYNAME LIKE '%\,%' ESCAPE '\';
```

### Commas in Book Titles
```
SELECT ZTITLE, ZDISPLAYNAME FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZTITLE LIKE '%\,%' ESCAPE '\';
```

### Parenthesis in Book Titles
We don't like book titles that say things like `(Book 4)` since that's not actually part of the title. Run the following query to find potential issues:
```
SELECT ZTITLE, ZDISPLAYNAME FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZTITLE LIKE '%\(%' ESCAPE '\';
```

### Missing Call Number
```
SELECT ZTITLE, ZDISPLAYNAME, ZLCC FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZLCC IS NULL;
```

### Empty Genres
```
SELECT ZTITLE, ZDISPLAYNAME, ZGENRE FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZGENRE LIKE "";
```

### Bad Genre Names
```
SELECT ZTITLE, ZDISPLAYNAME, ZGENRE FROM ZBOOK
INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
WHERE ZGENRE NOT LIKE "Art" AND
ZGENRE NOT LIKE "Biography & Autobiography" AND
ZGENRE NOT LIKE "Business & Economics" AND
ZGENRE NOT LIKE "Coloring Book" AND
ZGENRE NOT LIKE "Comics & Graphic Novels" AND
ZGENRE NOT LIKE "Cooking" AND 
ZGENRE NOT LIKE "Dictionaries" AND
ZGENRE NOT LIKE "Encyclopedias" AND
ZGENRE NOT LIKE "Fiction" AND 
ZGENRE NOT LIKE "Foreign Language Study" AND
ZGENRE NOT LIKE "Health & Fitness" AND
ZGENRE NOT LIKE "Humor" AND
ZGENRE NOT LIKE "Juvenile Fiction" AND 
ZGENRE NOT LIKE "Literary Criticism" AND
ZGENRE NOT LIKE "Magic" AND
ZGENRE NOT LIKE "Philosophy" AND
ZGENRE NOT LIKE "Poetry" AND
ZGENRE NOT LIKE "Religion" AND 
ZGENRE NOT LIKE "Self-Help" AND
ZGENRE NOT LIKE "Young Adult Fiction";
```
