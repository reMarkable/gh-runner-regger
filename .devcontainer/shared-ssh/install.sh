#!/usr/bin/env bash

set -e

# Setup .ssh
echo "Setting up shared ssh configuration..."
sudo mkdir -p "${_REMOTE_USER_HOME}/.ssh"
chown ${_REMOTE_USER}.${_REMOTE_USER} "${_REMOTE_USER_HOME}/.ssh"
chmod 0700 "${_REMOTE_USER_HOME}/.ssh"
echo "Done!"
