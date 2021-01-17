// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/s4kibs4mi/newschain-cache/ent/latestblock"
)

// LatestBlock is the model entity for the LatestBlock schema.
type LatestBlock struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber uint32 `json:"block_number"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LatestBlock) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case latestblock.FieldID, latestblock.FieldBlockNumber:
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type LatestBlock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LatestBlock fields.
func (lb *LatestBlock) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case latestblock.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			lb.ID = int(value.Int64)
		case latestblock.FieldBlockNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block_number", values[i])
			} else if value.Valid {
				lb.BlockNumber = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this LatestBlock.
// Note that you need to call LatestBlock.Unwrap() before calling this method if this LatestBlock
// was returned from a transaction, and the transaction was committed or rolled back.
func (lb *LatestBlock) Update() *LatestBlockUpdateOne {
	return (&LatestBlockClient{config: lb.config}).UpdateOne(lb)
}

// Unwrap unwraps the LatestBlock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lb *LatestBlock) Unwrap() *LatestBlock {
	tx, ok := lb.config.driver.(*txDriver)
	if !ok {
		panic("ent: LatestBlock is not a transactional entity")
	}
	lb.config.driver = tx.drv
	return lb
}

// String implements the fmt.Stringer.
func (lb *LatestBlock) String() string {
	var builder strings.Builder
	builder.WriteString("LatestBlock(")
	builder.WriteString(fmt.Sprintf("id=%v", lb.ID))
	builder.WriteString(", block_number=")
	builder.WriteString(fmt.Sprintf("%v", lb.BlockNumber))
	builder.WriteByte(')')
	return builder.String()
}

// LatestBlocks is a parsable slice of LatestBlock.
type LatestBlocks []*LatestBlock

func (lb LatestBlocks) config(cfg config) {
	for _i := range lb {
		lb[_i].config = cfg
	}
}
