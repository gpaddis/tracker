This program creates a record per day with a start time and an end time. Compile it and use the flags:

`./tracker --track`<br />
Inserts a new record for the current day in the database. If a record is already present, it updates the end time when the program is executed.

`./tracker --pause {time string}`<br />
Updates the pause duration for today (default: 60m). Use a time format like 1h10m, 30m, etc.

`./tracker --stats {time span}`<br />
Prints the statistics for the selected time span. Valid time spans are *today*, *yesterday*, *thisweek*, *lastweek*.