This program creates a record per day with a start time and an end time. Compile it and use the flags:

`./tracker --track`<br />
Inserts a new record with a start time in the database. If a record for that day is already present, it updates the end time when the program is executed.

`./tracker --stats timeSpan`<br />
Prints the statistics for the selected time span. Valid time spans are *today*, *yesterday*, *thisweek*, *lastweek*