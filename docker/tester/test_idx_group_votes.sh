#!/bin/bash

# TODO: group id
group_id=2

# test account address
address1=chora1l2pwmzk96ftmmt5egpjulyqtneygmmzns8r2ea
address2=chora14v5z5yyl5unnyu6q3ele8ze9jev6y0m7kg863q

# transaction flags
chora_tx_flags="--keyring-backend test --chain-id chora-local --yes"

# set group members json
cat > members.json <<EOL
{
  "members": [
    {
      "address": "$address1",
      "weight": "1",
      "metadata": ""
    },
    {
      "address": "$address2",
      "weight": "1",
      "metadata": ""
    }
  ]
}
EOL

# create test group
chora tx group create-group "$address1" "" members.json --from "$address1" $chora_tx_flags

# wait for transaction to be processed
sleep 10

# set group policy json
cat > policy.json <<EOL
{
  "@type": "/cosmos.group.v1.ThresholdDecisionPolicy",
  "threshold": "2",
  "windows": {
    "voting_period": "20s",
    "min_execution_period": "0s"
  }
}
EOL

# create test group policy
chora tx group create-group-policy "$address1" $group_id "" policy.json --from "$address1" $chora_tx_flags

# wait for transaction to be processed
sleep 10

# set group policy address
policy_address=$(chora q group group-policies-by-group $group_id --output json | jq -r '.group_policies[-1].address')

# set group proposal json
cat > proposal.json <<EOL
{
  "group_policy_address": "$policy_address",
  "messages": [],
  "metadata": "",
  "proposers": ["$address1"]
}
EOL

# create group proposal
chora tx group submit-proposal proposal.json --from "$address1" $chora_tx_flags

# wait for transaction to be processed
sleep 10

# set proposal id
proposal_id=$(chora q group proposals-by-group-policy "$policy_address" --output json | jq -r '.proposals[-1].id')

# vote "yes" on proposal with user 1
chora tx group vote "$proposal_id" "$address1" VOTE_OPTION_YES "" --from "$address1" $chora_tx_flags

# vote "yes" on proposal with user 2
chora tx group vote "$proposal_id" "$address2" VOTE_OPTION_YES "" --from "$address2" $chora_tx_flags

# wait for transactions to be processed
sleep 10

# TODO: confirm votes indexed in database
psql postgres://postgres:password@localhost:5432/postgres -c "SELECT * from idx_group_vote;"
# TODO: if votes NOT found, then exit 1
