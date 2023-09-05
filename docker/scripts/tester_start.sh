#!/bin/bash

set -eo pipefail

# wait for idx startup
sleep 5

# run tester test scripts
/home/tester/scripts/test_idx_group_proposals.sh
/home/tester/scripts/test_idx_group_votes.sh

# exit without error
exit 0
