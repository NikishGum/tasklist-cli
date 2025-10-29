
# CLI Task List Appliction

This projects represents simple __CLI Task List__ appliction using pure Golang, without any frameworks or external libraries

## Installation

Using *git clone*:

``
git clone https://github.com/NikishGum/tasklist-cli
``

## Usage

1. __Add task to be done__

> Assuming you already opened project folder in terminal

`` ./tasklist-cli add Some task``

This command will add a task to Task List with status *"To be done"* and description *"Some task"*. It will output ID, given to this specific task, in other commands task can be accessed by ID. 

2. __Delete task__

`` ./tasklist-cli delete ID ``

Where *ID* is integer value assigned to specific Task

3. __Change task status__

`` ./tasklist-cli mark-done ID ``

or 

`` ./tasklist-cli mark-in-progress ID ``

4. __List tasks__

You can either output saved tasks by their status: 

`` ./tasklist-cli list done `` 

`` ./tasklist-cli list in-progress ``

or output all known tasks

`` ./tasklist-cli list ``