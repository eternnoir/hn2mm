# hn2mm
Auto post Hack News top storys to MatterMost channel

## Usage

### Prepare Sqlite DB

1. Create Sqlite3 database `sqlite3 <databasepath>`
1. Create Table
```
CREATE TABLE `AL_POST_STORY` (
	`ID`	INTEGER,
	PRIMARY KEY(ID)
);
```

### Install hn2mm

* Install from source code.

```
$ go get https://github.com/eternnoir/hn2mm.git
$ hn2mm -c config.toml
```

* Config File `config.toml`
```
WebhookUrl = ""			// MatterMost channel webhook url
Channel = ""			// Channel name. If you need.
Username = ""			// Username. If you need.
CheckNewInterval = "5m" // Interval to check top stories on hacker news.
DbString = "/tmp/sqlite3.db" // Sqlite3 database path.
```
