# hashgardcli gov vote

## Description

Vote for an active proposal, options: Yes/No/NoWithVeto/Abstain


## Usage

```shell
hashgardcli gov vote [proposal-id] [option] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Vote for proposal

```shell
hashgardcli gov vote 1 yes --chain-id=hashgard --from $you_wallet_name
```

After the proposal is officially activated, you can vote on this proposal

 You can determine if you have entered the voting period by looking at the details of the proposal:

```txt
Committed at block 15135 (tx hash: F210C95024B4366513A411E2E1AA59C25B93CAB637B109293EC8EE2999E45D6C, response: {Code:0 Data:[] Log:Msg 0:  Info: GasWanted:200000 GasUsed:11404 Tags:[{Key:[97 99 116 105 111 110] Value:[118 111 116 101] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[118 111 116 101 114] Value:[103 97 114 100 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 102 52 119 119 101 116] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[112 114 111 112 111 115 97 108 45 105 100] Value:[1] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}] Codespace: XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0})

```

How to query vote

[query-vote](query-vote.md)

[query-votes](query-votes.md)

[tally](tally.md)
