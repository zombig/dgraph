package gql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMutationTxnBlock1(t *testing.T) {
	query := `
	query {
		me(func: eq(age, 34)) {
			uid
			friend {
				uid
				age
			}
		}
	}
`
	_, err := ParseMutation(query)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Invalid block: [query]")
}

func TestMutationTxnBlock2(t *testing.T) {
	query := `
	txn {
		query {
			me(func: eq(age, 34)) {
				uid
				friend {
					uid
					age
				}
			}
		}
	}
}
`
	_, err := ParseMutation(query)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Too many right curl")
}

// Is this okay?
//  - Doesn't contain mutation op inside txn block
//  - uid and age are in the same line
func TestMutationTxnBlock3(t *testing.T) {
	query := `txn{query{me(func:eq(age,34)){uidfriend{uid age}}}}}`
	_, err := ParseMutation(query)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Too many right curl")
}

func TestMutationTxnBlock4(t *testing.T) {
	query := `
	txn {
		mutation {
			set {
				"_:user1" <age> "45" .
			}
		}

		query {
			me(func: eq(age, 34)) {
				uid
				friend {
					uid
					age
				}
			}
		}
	}
`
	_, err := ParseMutation(query)
	require.Nil(t, err)
}
