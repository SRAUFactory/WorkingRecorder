# WorkingRecorder
WorkingRecorder is a simple task management and record working time.

## Usage
The tool is run with a command-line argument to specify the action:

*   **start**: Initiates a new work session. You will be prompted to enter what you are working on.
*   **stop**: Stops the current work session. You will be prompted to select an option:
    1.  Suspend work (Saves current progress, allows restarting later).
    2.  Continue to work on other tasks (Repeats the start process).
    3.  Finish today's work (Finalizes records for the day and reports).
*   **report**: Generates a summary report of total logged time per task category.
*   **now**: Displays the details of the currently logged work session (if one exists).

**Example:**
To start working:
`./WorkingRecorder start`

To generate a report:
`./WorkingRecorder report`

