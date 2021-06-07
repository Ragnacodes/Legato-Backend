# Tool Box
- Type = `tool_boxes`

### SubTypes & Data
Adding a tool box node to the scenario.
- `sleep` (sleeps for time seconds)
    - Data request to create:
    
        ```json
        {
            "data": {
                "time": 5
            }
        }
        ```
    
    - Data response
        ```json
        {
            "data": {
                "time": 5
            }
        }
        ```
  
- `repeater` (repeats the other remaining nodes in that path)
    - Data request to create:
        ```json
        {
            "data": {
                "count": 3
            }
        }
        ```
    
    - Data response
        ```json
        {
            "data": {
                "count": 3
            }
        }
        ```