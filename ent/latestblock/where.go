// Code generated by entc, DO NOT EDIT.

package latestblock

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/s4kibs4mi/newschain-cache/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BlockNumber applies equality check predicate on the "block_number" field. It's identical to BlockNumberEQ.
func BlockNumber(v uint32) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberEQ applies the EQ predicate on the "block_number" field.
func BlockNumberEQ(v uint32) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberNEQ applies the NEQ predicate on the "block_number" field.
func BlockNumberNEQ(v uint32) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberIn applies the In predicate on the "block_number" field.
func BlockNumberIn(vs ...uint32) predicate.LatestBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LatestBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBlockNumber), v...))
	})
}

// BlockNumberNotIn applies the NotIn predicate on the "block_number" field.
func BlockNumberNotIn(vs ...uint32) predicate.LatestBlock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LatestBlock(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBlockNumber), v...))
	})
}

// BlockNumberGT applies the GT predicate on the "block_number" field.
func BlockNumberGT(v uint32) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberGTE applies the GTE predicate on the "block_number" field.
func BlockNumberGTE(v uint32) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberLT applies the LT predicate on the "block_number" field.
func BlockNumberLT(v uint32) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberLTE applies the LTE predicate on the "block_number" field.
func BlockNumberLTE(v uint32) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBlockNumber), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LatestBlock) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LatestBlock) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LatestBlock) predicate.LatestBlock {
	return predicate.LatestBlock(func(s *sql.Selector) {
		p(s.Not())
	})
}
