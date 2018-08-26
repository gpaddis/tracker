This program creates a record per day with a start time and an end time. I've built it to keep track of my work hours, but you're free to use it for any tracking purpose.

Compile it and use the flags:

`./tracker --track`<br />
Inserts a new record for the current day in the database. If a record is already present, it updates the end time when the program is executed.

`./tracker --pause {time string}`<br />
Updates the pause duration for today (default: 60m). Use a time format like 1h10m, 30m, etc.

`./tracker --stats {time span}`<br />
Prints the statistics for the selected time span. Valid time spans are *today*, *yesterday*, *thisweek*, *lastweek*.

`./tracker --balance {days back}`<br />
Prints the total balance for the given number of days back (1 = today).

### Keep track of your work hours: set up a cronjob
To track your work hours, run the tracker with the flag `--track` every x minutes (to make sure you track everything, even if your pc crashes.) For instance, the cron schedule expression might look like [* 8-20 * * 1-5](https://crontab.guru/#*_8-20_*_*_1-5). The sqlite database will be created in your home directory (or any other directory you are running the binary from).