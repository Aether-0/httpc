#!/bin/bash

# Define variables for file paths
LOCAL_SCRIPT_FILE="httpc.py"
REMOTE_SCRIPT_URL="https://raw.githubusercontent.com/Aether-0/httpc/main/httpc.py"
INSTALL_PATH="/bin/httpc"

# Check if the local script file exists
if [ ! -f "$LOCAL_SCRIPT_FILE" ]; then
    echo "Local script file '$LOCAL_SCRIPT_FILE' not found. Attempting to download from $REMOTE_SCRIPT_URL..."
    
    # Download the script file from the remote URL
    wget -q "$REMOTE_SCRIPT_URL" -O "$LOCAL_SCRIPT_FILE"
    
    # Check if the download was successful
    if [ $? -ne 0 ]; then
        echo "Error: Failed to download the script from '$REMOTE_SCRIPT_URL'."
        exit 1
    fi
fi

# Check if sudo privileges are available
if [ "$(id -u)" -ne 0 ]; then
    echo "Error: This script requires superuser privileges to install httpc."
    exit 1
fi

# Write the contents of the script to the installation path
cat "$LOCAL_SCRIPT_FILE" | sudo tee "$INSTALL_PATH" > /dev/null

# Check if the installation was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to install httpc to '$INSTALL_PATH'."
    exit 1
fi

# Add execute permission to the installed script
sudo chmod +x "$INSTALL_PATH"

# Check if the permission change was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to set execute permission for '$INSTALL_PATH'."
    exit 1
fi

# Display success message
echo "httpc is now installed. You can run it using 'httpc' command."
