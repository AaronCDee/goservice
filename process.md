# Action

- expects values
- promises values
- defines an executed function which contains what the action does
- validates it's expectations
- validates it's promises
- can be rolled back

# Organizer

- Is called with an arbitrary set of values expected by it's actions
- Instantiates a context
- Instantiates each action and executes that action

# Context

- Is passed from one action to another
- Carries data between actions
- Can be failed
- Can be rolled back

Maybe the context is what contains all the actions...?