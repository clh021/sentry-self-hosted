#!/bin/bash

REPO_URL="https://github.com/getsentry/self-hosted"

# Clone the repository
git clone $REPO_URL
cd self-hosted

# Get the latest release tag
RELEASE_TAG=$(git describe --tags $(git rev-list --tags --max-count=1))

# Checkout the release tag
git checkout $RELEASE_TAG

# Optional: Display the current commit hash
CURRENT_COMMIT=$(git rev-parse HEAD)
echo "Repository is now at release tag: $RELEASE_TAG"
echo "Current commit: $CURRENT_COMMIT"
