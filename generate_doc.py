import re

def extract_commands_and_docs(file_path):
    """
    Extracts command names, documentation, and types from handler.go file.
    """
    
    commands = []
    with open(file_path, 'r') as f:
        content = f.read()

    register_command_pattern = re.compile(r'RegisterCommand\("([^"]+)",\s*([^,]+),\s*`([^`]+)`\)')
    matches = register_command_pattern.findall(content)

    for command, func_name, doc in matches:
        command_type = determine_command_type(func_name)
        commands.append({
            "name": command,
            "doc": doc.strip(),
            "type": command_type,
        })

    return commands

def determine_command_type(func_name):
    """
    Determines the command type based on the function name.

    Args:
        func_name (str): The name of the function associated with the command.

    Returns:
        str: The command type.
    """
    if func_name.startswith("h"):
        return "Hash"
    elif func_name.startswith("r") or func_name.startswith("l"):
        return "List"
    elif func_name == "Help":
        return "Help"
    else:
        return "String"

def generate_doc_go(commands, output_file_path):
    """
    Generates a doc.go file with the extracted commands, documentation, and types.

    Args:
        commands (list): A list of dictionaries containing command names, documentation, and types.
        output_file_path (str): The path to the output doc.go file.
    """
    with open(output_file_path, 'w') as f:
        f.write("/*\n")
        f.write("Package animus provides an in-memory database (similar to Redis) implemented in Go.\n\n")
        f.write("Animus supports various data types like strings, hashes, and lists, along with features\n")
        f.write("such as expiration handling (TTL, LRU), and basic commands.\n\n")
        f.write("Key Features:\n")

        for command in commands:
            f.write(f"  - **{command['name']} ({command['type']})**: {command['doc']}\n")

        f.write("\nRoadmap:\n")
        f.write("  - Advanced data structures (Sets, Sorted Sets)\n")
        f.write("  - Master-Slave replication\n")
        f.write("  - Pub/Sub for messaging\n")
        f.write("  - Performance optimizations\n")
        f.write("  - Clustering and sharding\n")
        f.write("*/\n")
        f.write("package main\n")

commands = extract_commands_and_docs("./internal/commands/handler.go")
generate_doc_go(commands, "doc.go")