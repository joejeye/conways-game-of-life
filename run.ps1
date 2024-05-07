# Get the path to the executable file
$executablePath = ".\exe\main.exe"

# Get the input arguments provided by the user
$arguments = $args

# Execute the executable file with the input arguments
& $executablePath $arguments
